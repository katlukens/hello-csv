package main

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {

	// Define your date range
	startDate, err := time.Parse("1/2/2006", "5/21/2024")
	if err != nil {
		slog.Error("Error parsing start date: %w", err)
	}

	endDate, err := time.Parse("1/2/2006", "6/5/2024")
	if err != nil {
		slog.Error("Error parsing end date: %w", err)
	}

	// List of files to parse
	files := []string{"AccountHistory.1.csv", "AccountHistory.2.csv", "AccountHistory.3.csv"}

	// Open a new CSV file for writing the results
	resultFile, err := os.Create("inDateRange.csv")
	if err != nil {
		fmt.Println("Error creating result file:", err)
		return
	}
	defer resultFile.Close()

	resultWriter := csv.NewWriter(resultFile)
	defer resultWriter.Flush()

	// Write the header to the result CSV file
	header := []string{"Account Number", "Post Date", "Check", "Description", "Debit", "Credit", "Status", "Balance"} // Adjust this header as needed
	if err := resultWriter.Write(header); err != nil {
		fmt.Println("Error writing header to result file:", err)
		return
	}

	// Loop through the files
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			print("Error opening file:", err)
			continue
		}

		// Create a CSV reader
		reader := csv.NewReader(f)

		// Read and discard the header row
		if _, err := reader.Read(); err != nil {
			fmt.Println("Error reading header:", err)
			return
		}

		// Read the CSV content
		for {
			record, err := reader.Read()
			if err != nil {
				break // Stop at EOF or other error
			}

			// Parse the date from the second column
			recordDate, err := time.Parse("1/2/2006", record[1])
			if err != nil {
				fmt.Println("Error parsing date:", err)
				continue // Skip rows with invalid dates
			}
			// Check if the date is within the range
			if recordDate.After(startDate) && recordDate.Before(endDate) {
				// Write the record to the result CSV file
				if err := resultWriter.Write(record); err != nil {
					fmt.Println("Error writing record to result file:", err)
				}
			}
		}
		f.Close()
	}
}
