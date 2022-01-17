package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
)

func main() {
	var output []rune
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	fmt.Print(url.QueryEscape(string(output[:len(output)-1])))
}
