// main.go
package main

import (
	"reflect"
	"tester/tester"
)

func test() string {
	return "Hello World!"
}

func main() {
	tester.TestCheckReturnType(reflect.TypeOf(test()).String(), "string", "has to be int")
}
