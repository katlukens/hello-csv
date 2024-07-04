package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func selectByDate() error {

	// Define your date range
	startDate, err := time.Parse("1/2/2006", "5/21/2024")
	if err != nil {
		return fmt.Errorf("unable to parse start date: %w", err)
	}

	endDate, err := time.Parse("1/2/2006", "6/5/2024")
	if err != nil {
		slog.Error("Error parsing end date: %w", err)
	}

	files := []string{"AccountHistory.1.csv", "AccountHistory.2.csv", "AccountHistory.3.csv"}

	// Open a new CSV file for writing the results
	resultFile, err := os.Create("inDateRange.csv")
	if err != nil {
		return fmt.Errorf("unable to create result file: %w", err)
	}
	defer resultFile.Close()

	resultWriter := csv.NewWriter(resultFile)
	defer resultWriter.Flush()

	// Write the header to the result CSV file
	header := []string{"Account Number", "Post Date", "Check", "Description", "Debit", "Credit", "Status", "Balance"} // Adjust this header as needed
	if err := resultWriter.Write(header); err != nil {
		return fmt.Errorf("unable to write header to result file: %w", err)
	}

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

			// Parse the date from the second column
			recordDate, err := time.Parse("1/2/2006", record[1])
			if err != nil {
				fmt.Println("unable to parse date:", err)
				continue // Skip rows with invalid dates
			}
			// Check if the date is within the range
			if recordDate.After(startDate) && recordDate.Before(endDate) {
				// Write the record to the result CSV file
				if err := resultWriter.Write(record); err != nil {
					fmt.Println("unable to write record to result file:", err)
				}
			}
		}
	}
	return nil
}
