package test

import (
	"encoding/json"
	"fmt"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/rest"
	"github.com/hem-nav-gateway/salesorder"
	"github.com/hem-nav-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SalesOrderResponseBody struct {
	Data   SalesOrderData       `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type SalesOrderData struct {
	SalesOrder       []salesorder.Model `json:"SalesOrder"`
	CreateSalesOrder salesorder.Model   `json:"CreateSalesOrder"`
	UpdateSalesOrder salesorder.Model   `json:"UpdateSalesOrder"`
}

func getAllSalesOrders() (int, SalesOrderResponseBody) {
	resBody := SalesOrderResponseBody{}
	page := utils.Query.SalesOrder
	attrs := utils.GetSalesOrderAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody

}

func createSalesOrder() (int, SalesOrderResponseBody, utils.ArgType) {
	//create salesorder
	resBody := SalesOrderResponseBody{}
	page := utils.Mutation.CreateSalesOrder
	attrs := utils.GetSalesOrderAttrs()
	args := utils.GetSalesOrderArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func updateSalesOrder() (int, SalesOrderResponseBody, utils.ArgType) {
	resBody := SalesOrderResponseBody{}
	page := utils.Mutation.UpdateSalesOrder
	attrs := utils.GetSalesOrderAttrs()
	args := utils.GetSalesOrderArgs().UpdateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func TestGetAllSalesOrder(t *testing.T) {
	resCode, resBody := getAllSalesOrders()
	elements := resBody.Data.SalesOrder
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")
}

func TestFilterSalesOrder(t *testing.T) {
	createSalesOrder()
	resBody := SalesOrderResponseBody{}
	page := utils.Query.SalesOrder
	attrs := utils.GetSalesOrderAttrs()
	args := utils.GetSalesOrderArgs().FilterArgs
	query := utils.GetQuery(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.SalesOrder
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
	responseCode, _ := rest.Delete(config.SalesOrderEndpoint, args, "Order")
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestCreateSalesOrder(t *testing.T) {
	resCode, resBody, args := createSalesOrder()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.CreateSalesOrder
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	// remove last element (Sales no need to check it)
	values = values[:len(values)-1]
	for _, val := range values {
		assert.NotNil(t, val)
	}
	responseCode, _ := rest.Delete(config.SalesOrderEndpoint, args, "Order")
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestUpdateSalesOrder(t *testing.T) {
	createSalesOrder()
	resCode, resBody, args := updateSalesOrder()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.UpdateSalesOrder
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	values = values[:len(values)-1]

	for _, val := range values {
		assert.NotNil(t, val)
	}
	responseCode, _ := rest.Delete(config.SalesOrderEndpoint, args, "Order")
	assert.Equal(t, 204, responseCode, "Could not delete entity")
}
