package utils_test

import (
	"DBWorker/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileContentsReturnsError(t *testing.T) {
	_, err := utils.GetFileContents("")
	assert.NotNil(t, err)
}
