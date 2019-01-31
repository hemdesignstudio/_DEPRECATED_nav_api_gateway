package test

import (
	"encoding/json"
	"github.com/nav-api-gateway/customer"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CustomerCardResponseBody struct {
	Data CustomerCardData `json:"data"`
}

type CustomerCardData struct {
	CustomerCard []customer.CustomerCard `json:"CustomerCard"`
}

func TestCustomerCardGetAll(t *testing.T) {
	resBody := CustomerCardResponseBody{}
	page := pages.customerCard
	attrs := getCustomerCardAttrs()
	query := getAllQuery(page, attrs)
	resCode, resBodyInBytes := client("GET", query)
	json.Unmarshal(resBodyInBytes, &resBody)
	element := resBody.Data.CustomerCard[0]

	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, element.No, "No should not be Nil")

}

func TestCustomerCardFilter(t *testing.T) {
	resBody := CustomerCardResponseBody{}
	page := pages.customerCard
	attrs := getCustomerCardAttrs()
	args := getCustomerCardArgs()
	queryList := getQueryList(page, attrs, args)

	for _, query := range queryList {
		resCode, resBodyInBytes := client("GET", query)
		json.Unmarshal(resBodyInBytes, &resBody)

		assert.Equal(t, 200, resCode, "Response code is 200 as expected")
		for _, element := range resBody.Data.CustomerCard {
			values := serialize(element)
			for _, val := range values {
				assert.NotNil(t, val)
			}
		}

	}
	assert.Equal(t, "12345", resBody.Data.CustomerCard[0].No, "Expected No = 12345")
}
