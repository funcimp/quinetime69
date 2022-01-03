package main

import (
	_ "embed"
	"fmt"
)

//go:embed head.txt
var head string

//go:embed body.txt
var body string

func main() {
	fmt.Printf("%s", head)
	for i, d := range body {
		if i == 0 {
			continue
		}
		fmt.Printf("\n\t%q,", string(d))
	}
	fmt.Printf("%s", body)
}
