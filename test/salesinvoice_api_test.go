package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
	"github.com/nav-api-gateway/salesinvoice"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SalesInvoiceResponseBody struct {
	Data   SalesInvoiceData     `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type SalesInvoiceData struct {
	SalesInvoice       []salesinvoice.SalesInvoice `json:"SalesInvoice"`
	CreateSalesInvoice salesinvoice.SalesInvoice   `json:"CreateSalesInvoice"`
	UpdateSalesInvoice salesinvoice.SalesInvoice   `json:"UpdateSalesInvoice"`
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

func createSalesInvoice() (int, SalesInvoiceResponseBody, utils.SliceOfMaps) {
	//create salesinvoice
	resBody := SalesInvoiceResponseBody{}
	page := utils.Mutation.CreateSalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	args := utils.GetSalesInvoiceArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func updateSalesInvoice() (int, SalesInvoiceResponseBody, utils.SliceOfMaps) {
	resBody := SalesInvoiceResponseBody{}
	page := utils.Mutation.UpdateSalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	args := utils.GetSalesInvoiceArgs().UpdateArgs
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
	createSalesInvoice()
	resBody := SalesInvoiceResponseBody{}
	page := utils.Query.SalesInvoice
	attrs := utils.GetSalesInvoiceAttrs()
	args := utils.GetSalesInvoiceArgs().FilterArgs
	queryList := utils.GetQueryList(page, attrs, args)
	for _, query := range queryList {
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
		navNo := args[0]["value"]
		assert.Equal(t, navNo, elements[0].No, fmt.Sprintf("Expected No = %s", navNo))
	}

	for _, arg := range args {
		arg["No"] = arg["value"]
		responseCode, _ := request.Delete(config.SalesInvoiceEndpoint, arg, "Invoice")
		assert.Equal(t, 204, responseCode, "Could not delete entity")
	}
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
	for _, arg := range args {
		responseCode, _ := request.Delete(config.SalesInvoiceEndpoint, arg, "Invoice")
		assert.Equal(t, 204, responseCode, "Could not delete entity")
	}
}

func TestUpdateSalesInvoice(t *testing.T) {
	createSalesInvoice()
	resCode, resBody, args := updateSalesInvoice()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.UpdateSalesInvoice
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	values = values[:len(values)-1]

	for _, val := range values {
		assert.NotNil(t, val)
	}
	for _, arg := range args {
		responseCode, _ := request.Delete(config.SalesInvoiceEndpoint, arg, "Invoice")
		assert.Equal(t, 204, responseCode, "Could not delete entity")
	}
}
