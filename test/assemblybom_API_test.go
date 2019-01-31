package test

import (
	"encoding/json"
	"fmt"
	"github.com/nav-api-gateway/assemblybom"
	"github.com/nav-api-gateway/roothandler"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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

type Resp struct {
	Data Data `json:"data"`
}

type Data struct {
	AssemblyBom []assemblybom.AssemblyBom `json:"assemblybom"`
}

func query() string {

	attrStr := strings.Join(attrs, " ")
	query := fmt.Sprintf(
		"?query={AssemblyBom(key:\"%s\",value:\"%s\"){%s}}",
		args[0]["key"],
		args[0]["value"],
		attrStr,
	)
	return query
}

func TestAssemblyBOMFiltering(t *testing.T) {
	res := Resp{}
	request, _ := http.NewRequest("GET", query(), nil)
	response := httptest.NewRecorder()
	handler := roothandler.RootEndpoint()
	handler.ServeHTTP(response, request)
	respBody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(respBody, &res)
	if err != nil {
		fmt.Println("could not unmarshal data")
	}
	assert.Equal(t, 200, response.Code, "Response code is 200 as expected")
	assert.Equal(t, "40001", res.Data.AssemblyBom[0].No, "Expected No = 40001")
	assert.Equal(t, "10005", res.Data.AssemblyBom[0].ParentItemNo, "Expected 10005")
	assert.Equal(t, "Item", res.Data.AssemblyBom[0].Type, "Expected Item")
	assert.Equal(t, 1.65, res.Data.AssemblyBom[0].QuantityPer, "Expected 1.65")
	assert.NotNil(t, res.Data.AssemblyBom[0].ParentItemNo)
}
