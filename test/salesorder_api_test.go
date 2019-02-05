package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/request"
	"github.com/nav-api-gateway/salesorder"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type SalesOrderResponseBody struct {
	Data   SalesOrderData       `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type SalesOrderData struct {
	SalesOrder       []salesorder.SalesOrder `json:"SalesOrder"`
	CreateSalesOrder salesorder.SalesOrder   `json:"CreateSalesOrder"`
	UpdateSalesOrder salesorder.SalesOrder   `json:"UpdateSalesOrder"`
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

func createSalesOrder() (int, SalesOrderResponseBody, utils.SliceOfMaps) {
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

func updateSalesOrder() (int, SalesOrderResponseBody, utils.SliceOfMaps) {
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
	queryList := utils.GetQueryList(page, attrs, args)
	for _, query := range queryList {
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
		navNo := args[0]["value"]
		assert.Equal(t, navNo, elements[0].No, fmt.Sprintf("Expected No = %s", navNo))
	}

	for _, arg := range args {
		arg["No"] = arg["value"]
		responseCode, _ := request.Delete(config.SalesOrderEndpoint, arg, "Order")
		assert.Equal(t, 204, responseCode, "Could not delete entity")
	}
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
	for _, arg := range args {
		responseCode, _ := request.Delete(config.SalesOrderEndpoint, arg, "Order")
		assert.Equal(t, 204, responseCode, "Could not delete entity")
	}
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
	for _, arg := range args {
		responseCode, _ := request.Delete(config.SalesOrderEndpoint, arg, "Order")
		assert.Equal(t, 204, responseCode, "Could not delete entity")
	}
}
