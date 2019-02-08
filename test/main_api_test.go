package test

import (
	"github.com/nav-api-gateway/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	resCode, _ := utils.Client("GET", "", nil)
	assert.Equal(t, 200, resCode, "Response code is 200 as expected")
}
