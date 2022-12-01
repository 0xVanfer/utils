package utils

import (
	"encoding/json"
	"io"
	"os"
)

// Read json file into []byte.
func ReadJsonBytes(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	return byteValue, err
}

// Read json file into []byte and convert into string.
func ReadJsonString(path string) (string, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	return string(byteValue), err
}

// Read json file into []byte and unmarshal into []string.
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

// Read json file into []byte and unmarshal into v.
func ReadJson(path string, v any) error {
	bytes, err := ReadJsonBytes(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	return err
}
