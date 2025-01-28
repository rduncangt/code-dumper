package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const outputFile = "code_dump.md"

// Directories to exclude
const excludedDirs = "logs:pkg"

// Supported file extensions and their markdown fencing language
var includedExtensions = map[string]string{
	".go":  "go",        // Go files
	".md":  "markdown",  // Markdown files
	".mod": "plaintext", // Go mod files
	".sum": "plaintext", // Go sum files
	".sql": "sql",       // SQL files
}

func main() {
	// Parse command-line flags
	allFiles := flag.Bool("all", false, "Include all supported non-binary files, not just Go files")
	flag.Parse()

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

		// Get file extension
		ext := filepath.Ext(path)

		// Process files based on the --all flag
		if *allFiles {
			// Include all supported file types
			if lang, ok := includedExtensions[ext]; ok {
				if err := processFile(path, lang, output); err != nil {
					return fmt.Errorf("error processing file %s: %v", path, err)
				}
			}
		} else {
			// Only include Go files by default
			if ext == ".go" {
				if err := processFile(path, "go", output); err != nil {
					return fmt.Errorf("error processing file %s: %v", path, err)
				}
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

// processFile reads the content of a file and writes it to the output file with fenced markdown
func processFile(path, lang string, output io.Writer) error {
	// Read the file content
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %s: %v", path, err)
	}

	// Write the file path and fenced code block to the output file
	_, err = fmt.Fprintf(output, "\n### %s\n\n```%s\n%s\n```\n", path, lang, string(content))
	if err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	return nil
}
