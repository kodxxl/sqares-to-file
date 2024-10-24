package main

import (
	"log/slog"
	"os"

	"github.com/kodxxl/sqares-to-file/calculator"
	"github.com/kodxxl/sqares-to-file/storage"
)

func main() {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sqares := calculator.GetSqaresOf(source)

	const FILENAME = "sqares.txt"

	err := storage.SaveToFile(FILENAME, sqares)
	if err != nil {
		slog.Info(err.Error())
		os.Exit(1)
	}
	slog.Info("Sqares successfully created")
}
