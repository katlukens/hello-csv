package main

import (
	"log/slog"
	"os"
)

func main() {
	err := selectByFood()
	if err != nil {
		slog.Error("Error selecting by food",
			"error", err,
		)
		os.Exit(1)
	}
	slog.Info("Successfully selected for food")
	os.Exit(0)
}
