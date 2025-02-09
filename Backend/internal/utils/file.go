package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load JSON from File
func LoadJSON(filePath string, target interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(target)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	return nil
}
