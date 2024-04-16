// main.go
package main

import (
	"tester/tester"
)

func main() {
	// Setting the value of RunTests variable from the tester package
	tester.RunTests = true
	tester.ShowOutput = true

	// Now you can call your testing functions
	tester.TestEqual(1, 2, "has to be equal")
	tester.PrintExecutedTests()
}
