// main.go
package main

import (
	"tester/tester"
)

func test() string {
	return "Hello World!"
}

func main() {
	tester.TestEqual(1, 1, "has to be equal")
}
