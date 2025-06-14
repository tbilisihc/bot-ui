package read_message

import (
	"fmt"
	"os"
)

func Read(filePath string) (string, error) {
	// os.ReadFile reads the entire file and returns its contents as a byte slice.
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		// If an error occurs, we wrap it with more context and return it.
		return "", fmt.Errorf("error reading file '%s': %w", filePath, err)
	}
	// Convert the byte slice to a string and return it.
	return string(contentBytes), nil
}
