package roothandler

import (
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/errorhandler"
)

func checkCompany(companyName string) bool {
	for key := range config.Companies {
		if key == companyName {
			return true
		}
	}
	return false
}

// pathVariables extracts vars from URL path and Returns Company
func pathVariables(vars map[string]string) (string, error) {
	name := vars["company"]

	if checkCompany(name) {
		return config.Companies[name], nil
	}
	return "", errorhandler.CompanyDoesNotExist(name)
}
