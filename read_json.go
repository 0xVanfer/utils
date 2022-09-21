package utils

import (
	"encoding/json"
	"io"
	"os"
)

func ReadJsonSlice(path string) ([]string, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var slice []string
	err = json.Unmarshal([]byte(byteValue), &slice)
	return slice, err
}
