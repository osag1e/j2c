package main

import (
	"encoding/json"
	"os"
)

func readJSONFile(jsonFilePath string) ([]map[string]interface{}, error) {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []map[string]interface{}

	var jsonData interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jsonData); err != nil {
		return nil, err
	}

	switch jsonData := jsonData.(type) {
	case map[string]interface{}:
		data = append(data, jsonData)
	case []interface{}:
		for _, item := range jsonData {
			if itemMap, ok := item.(map[string]interface{}); ok {
				data = append(data, itemMap)
			}
		}
	}
	return data, nil
}
