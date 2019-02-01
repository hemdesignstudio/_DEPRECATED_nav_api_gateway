package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func Serialize(element interface{}) []interface{} {

	elemObj := reflect.ValueOf(element)
	elemValues := make([]interface{}, elemObj.NumField())
	for i := 0; i < elemObj.NumField(); i++ {
		elemValues[i] = elemObj.Field(i).Interface()
	}
	return elemValues
}

func argsMapToString(args SliceOfMaps) string {
	var argStrList []string
	for _, element := range args {
		for key, value := range element {
			valueType := reflect.TypeOf(value).Kind()
			if valueType == reflect.String {
				argStrList = append(argStrList, fmt.Sprintf("%s:\"%s\"", key, value))
			} else {
				argStrList = append(argStrList, fmt.Sprintf("%s:%v", key, value))
			}
		}
	}

	argStr := strings.Join(argStrList, ",")
	return argStr
}

func GetAllQuery(page string, attrs []string) string {
	attrStr := strings.Join(attrs, " ")
	query := fmt.Sprintf(
		"?query={%s{%s}}",
		page,
		attrStr)
	return query
}

func GetQueryList(page string, attrs []string, args []map[string]interface{}) []string {
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

func GetPOSTBody(page string, attrs []string, args SliceOfMaps) map[string]interface{} {
	attrStr := strings.Join(attrs, " ")
	argStr := argsMapToString(args)
	body := map[string]interface{}{
		"query": fmt.Sprintf("mutation{%s(%s){%s}}", page, argStr, attrStr)}
	return body
}
