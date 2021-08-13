// an unsavory quine by funcimp.biz
package main

import "fmt"

func main() {
	your := make(chan func(), 1)
	go func() { your <- self }()
	(<-your)()
}

func self() {
	fmt.Printf("%s\n\nconst s = \x60%s\x60\n", s, s)
}

const s = `// an unsavory quine by funcimp.biz
package main

import "fmt"

func main() {
	your := make(chan func(), 1)
	go func() { your <- self }()
	(<-your)()
}

func self() {
	fmt.Printf("%s\n\nconst s = \x60%s\x60\n", s, s)
}`
