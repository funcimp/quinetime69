package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// this is used to generate the []string for the ken thompson quine
// to use: remove the contents of the var s and then run main.go through this formatter
// go run generator/main.go < main.go
// this generates properly formatted escaped single character strings for use in the quine

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		d, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		fmt.Printf("\t%q,\n", string(d))
	}
}
