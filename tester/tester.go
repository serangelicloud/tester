// tester.go
package tester

import (
	"fmt"
	"reflect"
)

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func TestEqual(actual interface{}, expected interface{}, request string) bool {
	if actual == expected {
		output := fmt.Sprintf("%sTest Passed%s - %v == %v", ColorGreen, ColorReset, actual, expected)
		fmt.Println(output)
		return true
	} else {
		output := fmt.Sprintf("%sTest Failed%s - %v != %v - %s", ColorRed, ColorReset, actual, expected, request)
		fmt.Println(output)
		return false
	}
}

func TestDifferent(actual interface{}, expected interface{}, request string) bool {
	if actual != expected {
		output := fmt.Sprintf("%sTest Passed%s - %v != %v", ColorGreen, ColorReset, actual, expected)
		fmt.Println(output)
		return true
	} else {
		output := fmt.Sprintf("%sTest Failed%s - %v == %v - %s", ColorRed, ColorReset, actual, expected, request)
		fmt.Println(output)
		return false
	}
}
func TestNil(actual interface{}, request string) bool {
	if actual == nil {
		output := fmt.Sprintf("%sTest Passed%s - %v == %v", ColorGreen, ColorReset, actual, nil)
		fmt.Println(output)
		return true
	} else {
		output := fmt.Sprintf("%sTest Failed%s - %v != %v - %s", ColorRed, ColorReset, actual, nil, request)
		fmt.Println(output)
		return true
	}
}
func TestNotNil(actual interface{}, request string) bool {
	if actual != nil {
		output := fmt.Sprintf("%sTest Passed%s - %v != %v", ColorGreen, ColorReset, actual, nil)
		fmt.Println(output)
		return true
	} else {
		output := fmt.Sprintf("%sTest Failed%s - %v == %v - %s", ColorRed, ColorReset, actual, nil, request)
		fmt.Println(output)
		return true
	}
}
func TestCheckReturnType(actual interface{}, expected interface{}, request string) bool {
	if reflect.TypeOf(actual).String() == expected {
		output := fmt.Sprintf("%sTest Passed%s - %v == %v", ColorGreen, ColorReset, reflect.TypeOf(actual).String(), expected)
		fmt.Println(output)
		return true
	} else {
		output := fmt.Sprintf("%sTest Failed%s - %v != %v - %s", ColorRed, ColorReset, reflect.TypeOf(actual).String(), expected, request)
		fmt.Println(output)
		return false
	}
}
