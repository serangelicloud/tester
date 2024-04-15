// main.go
package main

import (
	"tester/tester"
)

func test() string {
	return "Hello World!"
}

func main() {
	// Setting the value of RunTests variable from the tester package
	tester.RunTests = true

	// Now you can call your testing functions
	tester.TestEqual(1, 1, "has to be equal")
}
