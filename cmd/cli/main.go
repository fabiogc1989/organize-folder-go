package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fabiogc1989/organize-folder/internal/domain"
	"github.com/fabiogc1989/organize-folder/internal/usecase"
)

func main() {
	dirFLag := flag.String("dir", "", "Path to the folder to organize")
	flag.Parse()

	dirPath := *dirFLag
	if dirPath == "" {
		fmt.Fprintln(os.Stderr, "Error: -dir flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if !usecase.IsDirectoryPath(dirPath) {
		log.Fatalf("Invalid directory path: %s", dirPath)
	}

	rootDir, err := domain.NewDirectory(filepath.Base(dirPath), dirPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	policies, err := usecase.CreateDefaultPolicyStructsByPath(rootDir.Path)
	if err != nil {
		log.Fatal(err)
	}

	rootFiles, err := rootDir.GetFiles()
	if err != nil {
		log.Fatal(err)
	}

	var success, warning, failed int

	for _, fileInfo := range rootFiles {
		policy, policyFound := usecase.FindDirectoryByFileExtension(policies, fileInfo.Extension)
		if !policyFound {
			fmt.Printf("[Warning] No rules apply to %s, Size: %d bytes\n", fileInfo.Name, fileInfo.Size)
			warning++
			continue
		}

		if !policy.Dir.Exists() {
			if err := policy.Dir.Create(); err != nil {
				fmt.Printf("[Error] Failed to create directory %s: %v\n", policy.Dir.Name, err)
				failed++
				continue
			}
		}

		if err := fileInfo.Move(policy.Dir.Path); err != nil {
			fmt.Printf("[Error] Moving %s (Size: %d bytes): %v\n", fileInfo.Name, fileInfo.Size, err)
			failed++
			continue
		}

		fmt.Printf("[Success] %s moved to %s\n", fileInfo.Name, policy.Dir.Path)
		success++
	}

	fmt.Printf("\n\n%d item(s) moved\n%d item(s) not moved\n%d warning(s)\n", success, failed, warning)
}
