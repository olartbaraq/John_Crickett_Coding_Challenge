package tests

import (
	"testing"

	"github.com/olartbaraq/load_balancer/utils"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
}

func TestFileManager(t *testing.T) {
	var loadedData TestData
	err := utils.LoadFile("../files/test.json", &loadedData)
	assert.NoError(t, err)
	assert.FileExists(t, "../files/test.json")
}
