package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
	input, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("source open failed: %w", err)
	}
	defer input.Close()

	output, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("destination create failed: %w", err)
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		return fmt.Errorf("copy failed: %w", err)
	}
	return nil
}

func main() {
	err := copyFile("source.txt", "dest.txt")
	switch {
	case errors.Is(err, os.ErrNotExist):
		fmt.Println("Source file does not exist.")
	case errors.Is(err, os.ErrPermission):
		fmt.Println("Permission denied.")
	case errors.As(err, new(*os.PathError)):
		fmt.Println("Path error:", err)
	default:
		fmt.Println("Unexpected error:", err)
	}
}
