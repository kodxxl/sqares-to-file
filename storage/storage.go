package storage

import (
	"fmt"
	"os"
)

func SaveToFile(filename string, dataSource []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, value := range dataSource {
		fmt.Fprintln(file, value)
	}

	return nil
}
