package main

import (
	"fmt"
)

const head string = `package main

import "fmt"

var s = []string{`

const body string = `
}

func main() {
	fmt.Printf("package main\n\nimport \"fmt\"\n\nvar s = []string{\n")
	var o string
	for _, v := range s {
		fmt.Printf("\t%q,\n", v)
		o = o + v
	}
	fmt.Printf("%v", o)
}
`

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
