package helpers

import (
	"bufio"
	"log"
	"os"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func ReadFromConsole() string {
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error: something wrong during reading from console: %s", err.Error())
		return ""
	}
	return text
}
