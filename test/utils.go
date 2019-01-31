package test

import (
	"fmt"
	"reflect"
	"strings"
)

func serialize(element interface{}) []interface{} {

	elemObj := reflect.ValueOf(element)
	elemValues := make([]interface{}, elemObj.NumField())
	for i := 0; i < elemObj.NumField(); i++ {
		elemValues[i] = elemObj.Field(i).Interface()
	}
	return elemValues
}

func getAllQuery(page string, attrs []string) string {
	attrStr := strings.Join(attrs, " ")
	query := fmt.Sprintf(
		"?query={%s{%s}}",
		page,
		attrStr)
	return query
}

func getQueryList(page string, attrs []string, args []map[string]interface{}) []string {
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
