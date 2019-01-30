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
	"testing"
)

var args = []map[string]string{
	{
		"key":   "No",
		"value": "40001",
	},
}

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	AssemblyBom []assemblybom.AssemblyBom `json:"assemblybom"`
}

func query() string {

	query := fmt.Sprintf(
		"?query={AssemblyBom(key:\"%s\",value:\"%s\"){No Parent_Item_No}}",
		args[0]["key"],
		args[0]["value"],
	)
	return query
}

func TestAssemblyBOMFiltering(t *testing.T) {
	res := Response{}
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
	assert.NotNil(t, res.Data.AssemblyBom[0].ParentItemNo)
}
