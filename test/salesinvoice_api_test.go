package test

import (
	"encoding/json"
	"fmt"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/rest"
	"github.com/hem-nav-gateway/salesinvoice"
	"github.com/hem-nav-gateway/test/utils"

	"github.com/stretchr/testify/assert"
	"testing"
)

type SalesInvoiceResponseBody struct {
	Data   SalesInvoiceData     `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type SalesInvoiceData struct {
	SalesInvoice       []salesinvoice.Model `json:"SalesInvoice"`
	CreateSalesInvoice salesinvoice.Model   `json:"CreateSalesInvoice"`
	UpdateSalesInvoice salesinvoice.Model   `json:"UpdateSalesInvoice"`
}

func getAllSalesInvoices() (int, SalesInvoiceResponseBody) {
	resBody := SalesInvoiceResponseBody{}
	page := utils.Query.SalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody
}

func createSalesInvoice() (int, SalesInvoiceResponseBody, utils.ArgType) {
	resBody := SalesInvoiceResponseBody{}
	page := utils.Mutation.CreateSalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	args := utils.GetSalesInvoiceArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func updateSalesInvoice(id string) (int, SalesInvoiceResponseBody, utils.ArgType) {
	resBody := SalesInvoiceResponseBody{}
	page := utils.Mutation.UpdateSalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	args := utils.GetSalesInvoiceArgs().UpdateArgs
	args["No"] = id
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func TestGetAllSalesInvoice(t *testing.T) {
	resCode, resBody := getAllSalesInvoices()
	elements := resBody.Data.SalesInvoice
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")
}

func TestFilterSalesInvoice(t *testing.T) {
	_, resBody, _ := createSalesInvoice()
	page := utils.Query.SalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	args := utils.GetSalesInvoiceArgs().FilterArgs

	args["value"] = resBody.Data.CreateSalesInvoice.No

	query := utils.GetQuery(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.SalesInvoice
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements, "Empty results are returned")
	for _, element := range elements {
		values := utils.Serialize(element)
		values = values[:len(values)-1]
		for _, val := range values {
			assert.NotNil(t, val)
		}
	}
	navNo := args["value"]
	assert.Equal(t, navNo, elements[0].No, fmt.Sprintf("Expected No = %s", navNo))

	args["No"] = args["value"]
	responseCode, _ := rest.Delete(config.SalesInvoiceEndpoint, args, "Invoice")
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestCreateSalesInvoice(t *testing.T) {
	resCode, resBody, args := createSalesInvoice()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.CreateSalesInvoice
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	// remove last element (Sales no need to check it)
	values = values[:len(values)-1]
	for _, val := range values {
		assert.NotNil(t, val)
	}
	args["No"] = resBody.Data.CreateSalesInvoice.No
	responseCode, _ := rest.Delete(config.SalesInvoiceEndpoint, args, "Invoice")
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestUpdateSalesInvoice(t *testing.T) {
	_, createResBody, _ := createSalesInvoice()
	id := createResBody.Data.CreateSalesInvoice.No
	resCode, resBody, args := updateSalesInvoice(id)
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.UpdateSalesInvoice
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	values = values[:len(values)-1]

	for _, val := range values {
		assert.NotNil(t, val)
	}
	args["No"] = createResBody.Data.CreateSalesInvoice.No
	responseCode, _ := rest.Delete(config.SalesInvoiceEndpoint, args, "Invoice")
	assert.Equal(t, 204, responseCode, "Could not delete entity")
}
