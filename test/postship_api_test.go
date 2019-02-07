package test

import (
	"encoding/json"
	"github.com/nav-api-gateway/postship"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type PostShipResponseBody struct {
	Data   PostShipData         `json:"data"`
	Errors []utils.ErrorMessage `json:"errors"`
}

type PostShipData struct {
	PostShip []postship.PostShip `json:"PostedSalesShipment"`
}

func TestPostShipGetAll(t *testing.T) {
	resBody := PostShipResponseBody{}
	page := utils.Query.PostShip
	attrs := utils.GetPostShipAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	elements := resBody.Data.PostShip
	assert.NotEmpty(t, elements, "Empty results are returned")
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, elements[0].No, "No should not be Nil")
	assert.NotNil(t, elements[0].BillToAddress, "should not be Nil")
	assert.NotNil(t, elements[0].OrderNo, "should not be Nil")
	assert.NotNil(t, elements[0].PromisedDeliveryDate, "should not be Nil")
	assert.NotNil(t, elements[0].SellToCustomerNo, "should not be Nil")

}

func TestPostShipFilter(t *testing.T) {
	resBody := PostShipResponseBody{}
	page := utils.Query.PostShip
	attrs := utils.GetPostShipAttrs()
	args := utils.GetPostShipArgs().FilterArgs
	query := utils.GetQuery(page, attrs, args)

	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")

	elements := resBody.Data.PostShip
	assert.NotEmpty(t, elements, "Empty results are returned")
	for _, element := range elements {
		values := utils.Serialize(element)
		for _, val := range values {
			assert.NotNil(t, val)
		}

	}

	assert.Equal(t, "102005", resBody.Data.PostShip[0].No, "Expected No = 102005")
	assert.Equal(t, "WEB-10", resBody.Data.PostShip[0].BillToCustomerNo, "Expected 10005")
	assert.Equal(t, "MIERZYN", resBody.Data.PostShip[0].LocationCode, "Expected MIERZYN")
	assert.Equal(t, "115770", resBody.Data.PostShip[0].OrderNo, "Expected 115770")
}
