package test

import (
	"encoding/json"
	"github.com/hem-nav-gateway/inventory"
	"github.com/hem-nav-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type InventoryResponseBody struct {
	Data   InventoryData        `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type InventoryData struct {
	Inventory []inventory.Model `json:"inventory"`
}

func TestInventoryGetAll(t *testing.T) {
	resBody := InventoryResponseBody{}
	page := utils.Query.Inventory
	attrs := utils.GetInventoryAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.Inventory
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")
	assert.NotNil(t, elements[0].Description, "Description should not be Nil")
	assert.NotNil(t, elements[0].QuantityAvailable, "QuantityAvailable should not be Nil")
	assert.NotNil(t, elements[0].ReceiptDate, "ReceiptDate should not be Nil")

}

func TestInventoryFilter(t *testing.T) {
	resBody := InventoryResponseBody{}
	page := utils.Query.Inventory
	attrs := utils.GetInventoryAttrs()
	args := utils.GetInventoryArgs().FilterArgs
	query := utils.GetQuery(page, attrs, args)

	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")

	elements := resBody.Data.Inventory
	assert.NotEmpty(t, elements, "Empty results are returned")
	for _, element := range elements {
		values := utils.Serialize(element)
		for _, val := range values {
			assert.NotNil(t, val)
		}

	}

	assert.Equal(t, "10005", resBody.Data.Inventory[0].No, "Expected No = 10005")
	assert.Equal(t, "Hai Chair Mosaic Charcoal", resBody.Data.Inventory[0].Description, "Hai Chair Mosaic Charcoal")
	assert.Equal(t, "25.00", resBody.Data.Inventory[0].QuantityAvailable, "Expected 25.00")
}
