package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var page = "AssemblyBom"

var attrs = []string{
	"No",
	"Parent_Item_No",
	"No",
	"Type",
	"Quantity_per",
	"Unit_of_Measure_Code",
}

var args = []map[string]string{
	{
		"key":   "No",
		"value": "40001",
	},
}

type ResponseBody struct {
	Data Data `json:"data"`
}

type Data struct {
	AssemblyBom []assemblybom.AssemblyBom `json:"assemblybom"`
}

func getAllQuery() string {
	attrStr := strings.Join(attrs, " ")
	query := fmt.Sprintf(
		"?query={%s{%s}}",
		page,
		attrStr)
	return query
}

func getQueryList() []string {
	var queryList []string
	attrStr := strings.Join(attrs, " ")

	for _, element := range args {
		query := fmt.Sprintf(
			"?query={%s(key:\"%s\",value:\"%s\"){%s}}",
			page,
			element["key"],
			element["value"],
			attrStr)
		queryList = append(queryList, query)
	}

	return queryList
}

func TestGetAll(t *testing.T) {
	resBody := ResponseBody{}
	query := getAllQuery()
	resCode, resBodyInBytes := client("GET", query)
	json.Unmarshal(resBodyInBytes, &resBody)
	element := resBody.Data.AssemblyBom[0]

	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, element.No, "No should not be Nil")
	assert.NotNil(t, element.QuantityPer, "QuantityPer should not be Nil")
	assert.NotNil(t, element.Type, "Type should not be Nil")
	assert.NotNil(t, element.ParentItemNo, "ParentItemNo should not be Nil")
	assert.NotNil(t, element.UnitOfMeasureCode, "UnitOfMeasureCode should not be Nil")

}

func TestFilter(t *testing.T) {
	resBody := ResponseBody{}
	queryList := getQueryList()

	for _, query := range queryList {
		resCode, resBodyInBytes := client("GET", query)
		json.Unmarshal(resBodyInBytes, &resBody)

		assert.Equal(t, 200, resCode, "Response code is 200 as expected")

		for _, element := range resBody.Data.AssemblyBom {
			assert.NotNil(t, element.No, "No should not be Nil")
			assert.NotNil(t, element.QuantityPer, "QuantityPer should not be Nil")
			assert.NotNil(t, element.Type, "Type should not be Nil")
			assert.NotNil(t, element.ParentItemNo, "ParentItemNo should not be Nil")
			assert.NotNil(t, element.UnitOfMeasureCode, "UnitOfMeasureCode should not be Nil")

		}

	}

	assert.Equal(t, "40001", resBody.Data.AssemblyBom[0].No, "Expected No = 40001")
	assert.Equal(t, "10005", resBody.Data.AssemblyBom[0].ParentItemNo, "Expected 10005")
	assert.Equal(t, "Item", resBody.Data.AssemblyBom[0].Type, "Expected Item")
	assert.Equal(t, 1.65, resBody.Data.AssemblyBom[0].QuantityPer, "Expected 1.65")
}
