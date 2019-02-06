package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/item"
	"github.com/nav-api-gateway/request"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ItemCardResponseBody struct {
	Data   ItemCardData         `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type ItemCardData struct {
	ItemCard       []item.ItemCard `json:"ItemCard"`
	CreateItemCard item.ItemCard   `json:"CreateItemCard"`
	UpdateItemCard item.ItemCard   `json:"UpdateItemCard"`
}

func getAllItems() (int, ItemCardResponseBody) {
	resBody := ItemCardResponseBody{}
	page := utils.Query.ItemCard
	attrs := utils.GetItemCardAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody

}

func createItem() (int, ItemCardResponseBody, utils.ArgType) {
	//create item
	resBody := ItemCardResponseBody{}
	page := utils.Mutation.CreateItemCard
	attrs := utils.GetItemCardAttrs()
	args := utils.GetItemCardArgs().CreateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func updateItem() (int, ItemCardResponseBody, utils.ArgType) {
	resBody := ItemCardResponseBody{}
	page := utils.Mutation.UpdateItemCard
	attrs := utils.GetItemCardAttrs()
	args := utils.GetItemCardArgs().UpdateArgs
	body := utils.GetPOSTBody(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("POST", "", body)
	json.Unmarshal(resBodyInBytes, &resBody)
	return resCode, resBody, args
}

func TestGetAllItemCard(t *testing.T) {
	resCode, resBody := getAllItems()
	elements := resBody.Data.ItemCard
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")

}

func TestFilterItemCard(t *testing.T) {
	createItem()
	resBody := ItemCardResponseBody{}
	page := utils.Query.ItemCard
	attrs := utils.GetItemCardAttrs()
	args := utils.GetItemCardArgs().FilterArgs
	query := utils.GetQuery(page, attrs, args)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.ItemCard
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements, "Empty results are returned")
	for _, element := range elements {
		values := utils.Serialize(element)
		for _, val := range values {
			assert.NotNil(t, val)
		}

		navNo := args["value"]
		assert.Equal(t, navNo, elements[0].No, fmt.Sprintf("Expected No = %s", navNo))
	}

	args["No"] = args["value"]
	responseCode, _ := request.Delete(config.ItemCardEndpoint, args, nil)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestCreateItemCard(t *testing.T) {
	resCode, resBody, args := createItem()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.CreateItemCard
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	for _, val := range values {
		assert.NotNil(t, val)
	}
	responseCode, _ := request.Delete(config.ItemCardEndpoint, args, nil)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}

func TestUpdateItemCard(t *testing.T) {
	createItem()
	resCode, resBody, args := updateItem()
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	element := resBody.Data.UpdateItemCard
	assert.NotEmpty(t, element, "Empty results are returned")
	values := utils.Serialize(element)
	for _, val := range values {
		assert.NotNil(t, val)
	}
	responseCode, _ := request.Delete(config.ItemCardEndpoint, args, nil)
	assert.Equal(t, 204, responseCode, "Could not delete entity")

}
