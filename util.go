package adventofcode2023

import (
	"fmt"
	"os"
)

func ReadFile(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	return string(input)
}
