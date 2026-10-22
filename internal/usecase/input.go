package usecase

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadInput(r io.Reader, w io.Writer, isDirFunc func(string) bool) string {
	reader := bufio.NewReader(r)
	for {
		fmt.Fprintln(w, "Type a directory path to organize files:")

		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		input = strings.TrimSpace(input)

		if len(input) > 0 && isDirFunc(input) {
			return input
		}

		if len(input) == 0 {
			fmt.Println("Input cannot be empt. Please try again.")
		} else if isDirFunc(input) {
			fmt.Println("Invalid path. Please try again.")
		}
	}
}
