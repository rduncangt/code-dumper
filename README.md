# Code Dumper

A utility script to extract code from Go files in a project and format them into a Markdown file with fenced code blocks.

## Features

- Walks through the directory tree of your project.
- Extracts code from `.go` files.
- Outputs a Markdown file (`code_dump.md`) with file paths and fenced code blocks.

## Usage

Run the script in the root of your Go project to generate `code_dump.md`:

```bash
go run generate_code_dump.go
```

## Output Format

Each file is included with its path and a fenced code block:

### ./path/to/file.go

```go
// Code content
```

## Excluding Directories

Edit the `excludedDirs` constant to list directories you want to skip.

## License

[MIT License](./LICENSE)
