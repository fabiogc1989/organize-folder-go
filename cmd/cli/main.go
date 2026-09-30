package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fabiogc1989/organize-folder/internal/entity"
	"github.com/fabiogc1989/organize-folder/internal/usecase"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Type a directory path to organize files: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	dirPath := strings.TrimSpace(input)

	rootDir, err := entity.NewDirectory(filepath.Base(dirPath), dirPath, 0755)
	if err != nil {
		log.Fatalf("\033[31m%v\033[0m", err)
	}

	usecase.Directories, err = usecase.CreateDirectoryStructsByPath(rootDir.Path)
	if err != nil {
		log.Fatal(err)
	}

	if err = usecase.CreateDirectories(usecase.Directories); err != nil {
		log.Fatal(err)
	}

	rootFiles, err := rootDir.GetFiles()
	if err != nil {
		log.Fatalf("\033[31m%v\033[0m", err)
	}

	for _, fileInfo := range rootFiles {
		if strings.ToLower(fileInfo.Extension) == ".jpg" {
			pictureDir, pictureFound := usecase.FindDirectoryByName("Pictures")
			if pictureFound == false {
				log.Printf("\033[31mPictures directory not found: %v\033[0m", err)
				continue
			}

			if err = fileInfo.Move(pictureDir.Path); err != nil {
				log.Panic(err)
				continue
			}

			fmt.Printf("[MOVED] %s -> Pictures/\n", fileInfo.Name)
		} else {
			fmt.Printf("[FILE] %-15s Size: %d bytes\n", fileInfo.Name, fileInfo.Size)
		}
	}
}
