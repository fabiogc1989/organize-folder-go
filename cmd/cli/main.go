package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/fabiogc1989/organize-folder/internal/domain"
	"github.com/fabiogc1989/organize-folder/internal/usecase"
)

func main() {
	dirPath := usecase.ReadInput()

	rootDir, err := domain.NewDirectory(filepath.Base(dirPath), dirPath, 0755)
	if err != nil {
		log.Fatalf("\033[31m%v\033[0m", err)
	}

	policies, err := usecase.CreateDefaultPolicyStructsByPath(rootDir.Path)
	if err != nil {
		log.Fatal(err)
	}

	if err = usecase.CreateDirectories(policies); err != nil {
		log.Fatal(err)
	}

	rootFiles, err := rootDir.GetFiles()
	if err != nil {
		log.Fatalf("\033[31m%v\033[0m", err)
	}

	var (
		success int = 0
		warning int = 0
		error   int = 0
	)

	for _, fileInfo := range rootFiles {
		policy, policyFound := usecase.FindDirectoryByFileExtension(policies, fileInfo.Extension)
		if policyFound == false {
			fmt.Printf("[Warning] No rules apply to %s, Size: %d bytes\n", fileInfo.Name, fileInfo.Size)
			warning++
			continue
		}

		if err = fileInfo.Move(policy.Dir.Path); err != nil {
			fmt.Printf("[Error] Moving %s, Size: %d\n", fileInfo.Name, fileInfo.Size)
			error++
			continue
		}

		fmt.Printf("[Success] %s moved to %s\n", fileInfo.Name, policy.Dir.Path)
		success++
	}

	fmt.Printf("\n\n%d item(s) moved\n%d item(s) not moved\n%d warning(s)\n", success, error, warning)
}
