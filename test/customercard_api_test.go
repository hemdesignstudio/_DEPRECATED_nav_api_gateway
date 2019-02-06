package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/customer"
	"github.com/nav-api-gateway/request"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CustomerCardResponseBody struct {
	Data   CustomerCardData     `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type CustomerCardData struct {
	CustomerCard       []customer.CustomerCard `json:"CustomerCard"`
	CreateCustomerCard customer.CustomerCard   `json:"CreateCustomerCard"`
	UpdateCustomerCard customer.CustomerCard   `json:"UpdateCustomerCard"`
}

func getAllCustomers() (int, CustomerCardResponseBody) {
	resBody := CustomerCardResponseBody{}
	page := utils.Query.CustomerCard
	attrs := utils.GetCustomerCardAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody

}

func createCustomer() (int, CustomerCardResponseBody, utils.ArgType) {
	//create customer
	resBody := CustomerCardResponseBody{}
	page := utils.Mutation.CreateCustomerCard
	attrs := utils.GetCustomerCardAttrs()
	args := utils.GetCustomerCardArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func updateCustomer() (int, CustomerCardResponseBody, utils.ArgType) {
	resBody := CustomerCardResponseBody{}
	page := utils.Mutation.UpdateCustomerCard
	attrs := utils.GetCustomerCardAttrs()
	args := utils.GetCustomerCardArgs().UpdateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func TestGetAllCustomerCard(t *testing.T) {
	resCode, resBody := getAllCustomers()
	elements := resBody.Data.CustomerCard
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")

}

func TestFilterCustomerCard(t *testing.T) {
	createCustomer()
	resBody := CustomerCardResponseBody{}
	page := utils.Query.CustomerCard
	attrs := utils.GetCustomerCardAttrs()
	args := utils.GetCustomerCardArgs().FilterArgs
	query := utils.GetQuery(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.CustomerCard
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotEmpty(t, elements, "Empty results are returned")
	for _, element := range elements {
		values := utils.Serialize(element)
		for _, val := range values {
			assert.NotNil(t, val)
		}
	}
	navNo := args["value"]
	assert.Equal(t, navNo, elements[0].No, fmt.Sprintf("Expected No = %s", navNo))

	args["No"] = args["value"]
	responseCode, _ := request.Delete(config.CustomerCardWSEndpoint, args, nil)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestCreateCustomerCard(t *testing.T) {
	resCode, resBody, args := createCustomer()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.CreateCustomerCard
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	for _, val := range values {
		assert.NotNil(t, val)
	}
	responseCode, _ := request.Delete(config.CustomerCardWSEndpoint, args, nil)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestUpdateCustomerCard(t *testing.T) {
	createCustomer()
	resCode, resBody, args := updateCustomer()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.UpdateCustomerCard
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	for _, val := range values {
		assert.NotNil(t, val)
	}
	responseCode, _ := request.Delete(config.CustomerCardWSEndpoint, args, nil)
	assert.Equal(t, 204, responseCode, "Could not delete entity")
}
