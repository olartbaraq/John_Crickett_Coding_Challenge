package tests

import (
	"testing"

	"github.com/olartbaraq/load_balancer/utils"
)

type TestData struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
}

func TestFileManager(t *testing.T) {
	var loadedData TestData
	err := utils.LoadFile("../test.json", &loadedData)

}
