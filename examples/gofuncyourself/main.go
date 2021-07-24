package main

import "fmt"

func main() {
	your := make(chan func(), 1)
	go func() { your <- self }()
	(<-your)()
}

func self() {
	fmt.Printf("%v\n\nconst program = \u0060%v\u0060\n", program, program)
}

const program = `package main

import "fmt"

func main() {
	your := make(chan func(), 1)
	go func() { your <- self }()
	(<-your)()
}

func self() {
	fmt.Printf("%v\n\nconst program = \u0060%v\u0060\n", program, program)
}`
