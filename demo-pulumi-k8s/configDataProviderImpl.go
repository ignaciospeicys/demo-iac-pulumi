package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JSONConfigDataProvider struct {
	filePath string
}

func NewJSONConfigDataProvider(filePath string) *JSONConfigDataProvider {
	return &JSONConfigDataProvider{
		filePath: filePath,
	}
}

func (provider *JSONConfigDataProvider) GetConfigData() (map[string]string, error) {
	// Open the JSON file
	file, err := os.Open(provider.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Define a map to hold the configuration data
	var configData map[string]string

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(byteValue, &configData)
	if err != nil {
		return nil, err
	}

	// Return the configuration data map
	return configData, nil
}
