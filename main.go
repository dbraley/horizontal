package main

import (
	"bufio"
	"fmt"
	"github.com/dbraley/horizontal/lib"
	"io"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo '{\"level\":\"info\",\"time\":0,\"message\":\"structured logging output\"}' | horizontal")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	writer := lib.ConsoleWriter{
		Out:     os.Stdout,
		NoColor: false,
	}

	for {
		lineBytes, err := reader.ReadBytes('\n')

		if err != nil && err == io.EOF {
			break
		}

		// Trim trailing newline
		if len(lineBytes) > 0 {
			lineBytes = lineBytes[:len(lineBytes)-1]
		}

		_, err = writer.Write(lineBytes)

		// If we couldn't format the line, just print it
		if err != nil {
			fmt.Printf("!!! %s !!!\n", string(lineBytes))
		}
	}
}
