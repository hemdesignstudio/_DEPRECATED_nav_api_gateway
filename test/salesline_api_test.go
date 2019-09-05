package test

import (
	"encoding/json"
	"fmt"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/rest"
	"github.com/hem-nav-gateway/salesline"
	"github.com/hem-nav-gateway/test/utils"

	"github.com/stretchr/testify/assert"
	"testing"
)

type SalesLineResponseBody struct {
	Data   SalesLineData        `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type SalesLineData struct {
	SalesLine       []salesline.Model `json:"SalesLine"`
	CreateSalesLine salesline.Model   `json:"CreateSalesLine"`
	UpdateSalesLine salesline.Model   `json:"UpdateSalesLine"`
}

func getAllSalesLines() (int, SalesLineResponseBody) {
	resBody := SalesLineResponseBody{}
	page := utils.Query.SalesLine
	attrs := utils.GetSalesLineAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody
}

func createSalesLine() (int, SalesLineResponseBody, utils.ArgType) {
	//create salesline
	resBody := SalesLineResponseBody{}
	page := utils.Mutation.CreateSalesLine
	attrs := utils.GetSalesLineAttrs()
	args := utils.GetSalesLineArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func updateSalesLine() (int, SalesLineResponseBody, utils.ArgType) {
	resBody := SalesLineResponseBody{}
	page := utils.Mutation.UpdateSalesLine
	attrs := utils.GetSalesLineAttrs()
	args := utils.GetSalesLineArgs().UpdateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func TestGetAllSalesLine(t *testing.T) {
	resCode, resBody := getAllSalesLines()
	elements := resBody.Data.SalesLine
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")
	assert.NotNil(t, elements[0].Type, "No should not be Nil")
	assert.NotNil(t, elements[0].UnitOfMeasureCode, "No should not be Nil")
	assert.NotNil(t, elements[0].Description, "No should not be Nil")

}

func TestFilterSalesLine(t *testing.T) {
	_, resBody, _ := createSalesLine()
	page := utils.Query.SalesLine
	attrs := utils.GetSalesLineAttrs()
	args := utils.GetSalesLineArgs().FilterArgs
	args["value"] = resBody.Data.CreateSalesLine.No
	query := utils.GetQuery(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.SalesLine
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
	docType := resBody.Data.CreateSalesLine.DocumentType
	args["Document_No"] = resBody.Data.CreateSalesLine.DocumentNo
	args["Line_No"] = resBody.Data.CreateSalesLine.LineNo
	responseCode, _ := rest.Delete(config.SalesLineEndpoint, args, docType)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestCreateSalesLine(t *testing.T) {
	resCode, resBody, args := createSalesLine()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.CreateSalesLine
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	// remove last element (Sales no need to check it)
	for _, val := range values {
		assert.NotNil(t, val)
	}
	docType := args["Document_Type"]
	responseCode, _ := rest.Delete(config.SalesLineEndpoint, args, docType)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestUpdateSalesLine(t *testing.T) {
	createSalesLine()
	resCode, resBody, args := updateSalesLine()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.UpdateSalesLine
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)

	for _, val := range values {
		assert.NotNil(t, val)
	}
	docType := args["Document_Type"]
	responseCode, _ := rest.Delete(config.SalesLineEndpoint, args, docType)
	assert.Equal(t, 204, responseCode, "Could not delete entity")
}
