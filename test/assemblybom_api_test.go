package test

import (
	"encoding/json"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type AssemblyBomResponseBody struct {
	Data AssemblyBomData `json:"data"`
}

type AssemblyBomData struct {
	AssemblyBom []assemblybom.AssemblyBom `json:"assemblybom"`
}

func TestAssemblyBomGetAll(t *testing.T) {
	resBody := AssemblyBomResponseBody{}
	page := utils.Query.AssemblyBom
	attrs := utils.GetAssemblyBomAttrs()
	query := utils.GetAllQuery(page, attrs)
	resCode, resBodyInBytes := utils.Client("GET", query, nil)
	json.Unmarshal(resBodyInBytes, &resBody)
	element := resBody.Data.AssemblyBom[0]

	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
	assert.NotNil(t, element.No, "No should not be Nil")
	assert.NotNil(t, element.QuantityPer, "QuantityPer should not be Nil")
	assert.NotNil(t, element.Type, "Type should not be Nil")
	assert.NotNil(t, element.ParentItemNo, "ParentItemNo should not be Nil")
	assert.NotNil(t, element.UnitOfMeasureCode, "UnitOfMeasureCode should not be Nil")

}

func TestAssemblyBomFilter(t *testing.T) {
	resBody := AssemblyBomResponseBody{}
	page := utils.Query.AssemblyBom
	attrs := utils.GetAssemblyBomAttrs()
	args := utils.GetAssemblyBomArgs()
	queryList := utils.GetQueryList(page, attrs, args)

	for _, query := range queryList {
		resCode, resBodyInBytes := utils.Client("GET", query, nil)
		json.Unmarshal(resBodyInBytes, &resBody)

		assert.Equal(t, 200, resCode, "Response code is 200 as expected")
		elements := resBody.Data.AssemblyBom
		if len(elements) < 1 {
			assert.Equal(t, true, false, "Elements are not created, thus test fails")
		}
		for _, element := range elements {
			values := utils.Serialize(element)
			for _, val := range values {
				assert.NotNil(t, val)
			}

		}

	}

	assert.Equal(t, "40001", resBody.Data.AssemblyBom[0].No, "Expected No = 40001")
	assert.Equal(t, "10005", resBody.Data.AssemblyBom[0].ParentItemNo, "Expected 10005")
	assert.Equal(t, "Item", resBody.Data.AssemblyBom[0].Type, "Expected Item")
	assert.Equal(t, 1.65, resBody.Data.AssemblyBom[0].QuantityPer, "Expected 1.65")
}
