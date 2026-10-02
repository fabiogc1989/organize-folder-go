package usecase

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput() string {
	var (
		input string
		err   error
	)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Type a directory path to organize files:")

		input, err = reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}

		// Trim \r and \n (works across Windows and Linux/macOS)
		input = strings.TrimSpace(input)

		if len(input) > 0 && IsDirectoryPath(input) == true {
			return input
		}

		if len(input) == 0 {
			fmt.Println("Input cannot be empt. Please try again.")
		} else if IsDirectoryPath(input) == false {
			fmt.Println("Invalid path. Please try again.")
		}
	}
}
