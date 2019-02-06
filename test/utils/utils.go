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

func argsToString(args ArgType) string {
	var argStrList []string
	for key, value := range args {
		valueType := reflect.TypeOf(value).Kind()
		if valueType == reflect.String {
			argStrList = append(argStrList, fmt.Sprintf("%s:\"%s\"", key, value))
		} else {
			argStrList = append(argStrList, fmt.Sprintf("%s:%v", key, value))
		}

	}

	argStr := strings.Join(argStrList, ",")
	return argStr
}

func GetAllQuery(page string, attrs Attr) string {
	attrStr := strings.Join(attrs, " ")
	query := fmt.Sprintf(
		"?query={%s{%s}}",
		page,
		attrStr)
	return query
}

func GetQuery(page string, attrs Attr, args ArgType) string {
	attrStr := strings.Join(attrs, " ")
	query := fmt.Sprintf(
		"?query={%s(key:\"%s\",value:\"%s\"){%s}}",
		page,
		args["key"],
		args["value"],
		attrStr)
	return query
}

func GetPOSTBody(page string, attrs Attr, args ArgType) map[string]interface{} {
	attrStr := strings.Join(attrs, " ")
	argStr := argsToString(args)
	body := map[string]interface{}{
		"query": fmt.Sprintf("mutation{%s(%s){%s}}", page, argStr, attrStr)}
	return body
}
