package utils_test

import (
	"DBWorker/utils"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetFileContentsReturnsError(t *testing.T) {
	_, err := utils.GetFileContents("")
	assert.NotNil(t, err)
}

func TestGetAllTokensReturnsSuccess(t *testing.T) {
	src := "prefix :token12 suffix "
	regex, err := regexp.Compile(":[a-zA-Z][a-zA-Z0-9]+")
	assert.Nil(t, err, "Expected Regex compile to pass")

	noOfTokens, tokens := utils.GetAllTokens(src, regex)

	assert.Equal(t, 1, noOfTokens, "expected one token")
	assert.Equal(t, "token12", tokens[0], "expected token to match")
}

func TestGetAllTokensReturnsNoTokenForIncorrectTokens(t *testing.T) {
	src := "prefix :1token12 :9 suffix "
	regex, err := regexp.Compile(":[a-zA-Z][a-zA-Z0-9]+")
	assert.Nil(t, err, "Expected Regex compile to pass")

	noOfTokens, _ := utils.GetAllTokens(src, regex)

	assert.Equal(t, 0, noOfTokens, "expected one token")
}
