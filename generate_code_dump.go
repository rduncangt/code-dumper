package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const outputFile = "code_dump.md"
const excludedDirs = "logs:pkg" // Add directories to exclude, separated by ':'

func main() {
	// Open the output markdown file
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer output.Close()

	// Get the current working directory
	root, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Walk through the project directory
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing %s: %v", path, err)
		}

		// Skip directories that match excludedDirs
		if info.IsDir() && strings.Contains(excludedDirs, filepath.Base(path)) {
			return filepath.SkipDir
		}

		// Process only Go files
		if filepath.Ext(path) == ".go" {
			if err := processFile(path, output); err != nil {
				return fmt.Errorf("error processing file %s: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through project: %v\n", err)
		return
	}

	fmt.Println("Code dump generated:", outputFile)
}

// processFile reads the content of a Go file and writes it to the output file with fenced markdown
func processFile(path string, output io.Writer) error {
	// Read the file content
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", path, err)
	}

	// Write the file path and fenced code block to the output file
	_, err = fmt.Fprintf(output, "\n### %s\n\n```go\n%s\n```\n", path, string(content))
	if err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	return nil
}
