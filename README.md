# Code Dumper

A utility script to extract code from Go files in a project and optionally include other relevant non-binary files, formatting them into a Markdown file with fenced code blocks.

## Features

- Walks through the directory tree of your project.
- Extracts code from `.go` files by default.
- Optionally includes files like `.md`, `.mod`, `.sql`, and `.sum` with the `--all` flag.
- Outputs a Markdown file (`code_dump.md`) with file paths and fenced code blocks.

## Usage

Run the script in the root of your Go project to generate `code_dump.md`:

### Default: Process only Go files

```bash
go run generate_code_dump.go
```

### Extended: Process all supported files

```bash
go run generate_code_dump.go --all
```

## Output Format

Each file is included with its path and a fenced code block:
Example for Go File

### ./path/to/file.go

```go
// Code content


### Example for Markdown File
```markdown
### ./path/to/file.md

```markdown
# Markdown Title

Some content here.


## Supported File Types

By default, only `.go` files are processed. With the `--all` flag, the following file types are included:
- `.go` → `go`
- `.md` → `markdown`
- `.mod`, `.sum` → `plaintext`
- `.sql` → `sql`

## Excluding Directories

Edit the `excludedDirs` constant in the script to list directories you want to skip.

## License

[MIT License](./LICENSE)
