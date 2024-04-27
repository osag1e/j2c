package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func writeCSVToFile(data []map[string]interface{}, csvFilePath string) error {
	file, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(data) == 0 {
		return fmt.Errorf("no data to write")
	}

	var headers []string
	for key := range data[0] {
		headers = append(headers, key)
	}
	if err = writer.Write(headers); err != nil {
		return err
	}

	for _, record := range data {
		var row []string
		for _, key := range headers {
			value := fmt.Sprintf("%v", record[key])
			row = append(row, value)
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
