package utils

import (
	"encoding/json"
	"os"
)

func LoadFile[T any](filePath string, loadInto *T) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, loadInto)
	if err != nil {
		return err
	}
	return nil
}
