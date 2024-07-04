package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func selectByFood() error {

	// name file to be read
	files := []string{"inDateRangeFood.csv"}

	// Open a new CSV file for writing the results
	resultFile, err := os.Create("resultFood.csv")
	if err != nil {
		return fmt.Errorf("unable to create result file: %w", err)
	}

	defer resultFile.Close()

	// this writes the outfile
	resultWriter := csv.NewWriter(resultFile)
	defer resultWriter.Flush()

	// Write the header to the result CSV file
	header := []string{"Post Date", "Description", "Debit", "Status"}

	if err := resultWriter.Write(header); err != nil {
		return fmt.Errorf("unable to write header to result file: %w", err)
	}

	columnsToKeep := []int{1, 3, 4, 6}

	// Loop through the files even though there is only one
	// This is because I didn't want to change the code too much
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			print("Error opening file:", err) // Change this too?
			continue
		}

		defer f.Close()

		reader := csv.NewReader(f)

		// Read and discard the header row
		if _, err := reader.Read(); err != nil {
			return fmt.Errorf("unable to read header: %w", err)
		}

		for {
			record, err := reader.Read()
			if err != nil {
				break // Stop at EOF or other error
			}

			// parse the bool from the 9th column
			if record[8] == "1" {
				// Select the columns to keep
				selectedColumns := make([]string, len(columnsToKeep))
				for i, colIndex := range columnsToKeep {
					selectedColumns[i] = record[colIndex]
				}

				// Write the record to the result CSV file
				if err := resultWriter.Write(selectedColumns); err != nil {
					return fmt.Errorf("unable to write record to result file: %w", err)
				}
			}
		}
	}
	return nil
}
