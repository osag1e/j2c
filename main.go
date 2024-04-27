package main

import "fmt"

const (
	greenColour  = "\033[1;32m"
	purpleColour = "\033[1;34m"
	resetColour  = "\033[0m"
)

func main() {
	j2cArt := `
       _____   ______
      / /__ \ / ____/
 __  / /__/ // /     
/ /_/ // __// /___   
\____//____/\____/                      
`
	fmt.Printf("%s%sEnter JSON file path:%s\n", greenColour, j2cArt, resetColour)
	var jsonFilePath string
	_, err := fmt.Scanln(&jsonFilePath)
	if err != nil {
		panic("Error reading JSON file path: " + err.Error())
	}

	fmt.Printf("%sEnter CSV file path:%s\n", greenColour, resetColour)
	var csvFilePath string
	_, err = fmt.Scanln(&csvFilePath)
	if err != nil {
		panic("Error reading CSV file path: " + err.Error())
	}

	data, err := readJSONFile(jsonFilePath)
	if err != nil {
		panic("Error reading JSON file: " + err.Error())
	}

	err = writeCSVToFile(data, csvFilePath)
	if err != nil {
		panic("Error writing CSV file: " + err.Error())
	}

	fmt.Printf("%sJSON to CSV conversion complete.\nCSV file saved at: %s%s\n", purpleColour, csvFilePath, resetColour)
}
