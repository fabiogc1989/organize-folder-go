# organize-folder

A simple Go CLI that organizes files in a target directory into folders named after their file extensions, built with a modular DDD architecture.

## Overview

`organize-folder` takes a directory path as input, inspects the files inside it, and moves each file into a subdirectory corresponding to its extension. The codebase is structured using Domain-Driven Design (DDD) principles to keep domain rules decoupled from CLI input and filesystem operations.

## Project Structure

```text
.
├── cmd/
│   └── cli/
│       └── main.go           # CLI entrypoint
├── internal/
│   ├── domain/               # Core entities, file extension rules, and domain logic
│   └── usecase/              # Orchestrates the application workflow—reading directory contents, applying domain rules to determine target folders, and executing file moves
├── go.mod
├── README.md
└── LICENSE
````

## Installation

### Prerequisites

- Go 1.22 or higher

### Build from Source

```bash
git clone https://github.com/fabiogc1989/organize-folder-go.git
cd organize-folder-go
go build -o dist/organize-folder ./cmd/cli
```
(On Windows, use `dist/organize-folder.exe` as the output filename.)

## Usage

Pass the target directory path to the CLI:
```bash
./dist/organize-folder -dir /path/to/directory
```

On Windows (PowerShell):
```powershell
.\dist\organize-folder.exe -dir C:\path\to\directory
```

## Example

Before:
```text
Downloads/
├── invoice.docx
├── report.docx
└── picture1.jpeg
```

After:
```text
Downloads/
├── Documents/
│   ├── invoice.docx
│   └── report.docx
├── Pictures/
│   └── picture1.jpg
```
