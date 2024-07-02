package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func SelectByFood() {
	inFile, err := os.Open("inDateRangeFoo.csv")
	if err != nil {
		fmt.Println("Error opening file: %w", err)
		return
	}
	defer inFile.Close()

	reader := csv.NewReader(inFile)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data: %w", err)
		return
	}

	outFile, err := os.Create("filteredFood.csv")
	if err != nil {
		fmt.Println("Error creating file: %w", err)
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Iterate through the records, filtering and selecting columns
	for _, record := range records {
		if record[8] == "1" {
			selectedColumns := []string{record[1], record[3], record[4]}
			if err := writer.Write(selectedColumns); err != nil {
				fmt.Println("Error writing record to CSV: %w", err)
				return
			}
		}
	}
}
