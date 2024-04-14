// tester.go
package tester

import (
	"fmt"
	"reflect"
)

type output struct {
	TestType      string
	TestPosition  int
	ActualValue   interface{}
	ExpectedValue interface{}
}

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
		result := output{
			TestType:      "test for Equality",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: expected,
		}
		OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	} else {
		result := output{
			TestType:      "test for Equality",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: expected,
		}
		OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return false
	}
}

func TestDifferent(actual interface{}, expected interface{}, request string) bool {
	if actual != expected {
		result := output{
			TestType:      "test for difference",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: expected,
		}
		OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	} else {
		result := output{
			TestType:      "test for differences",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: expected,
		}
		OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return false
	}
}
func TestNil(actual interface{}, request string) bool {
	if actual == nil {
		result := output{
			TestType:      "test for Nil return",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: nil,
		}
		OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	} else {
		result := output{
			TestType:      "test for nil return",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: nil,
		}
		OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	}
}
func TestNotNil(actual interface{}, request string) bool {
	if actual != nil {
		result := output{
			TestType:      "test for Equality",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: "not nil",
		}
		OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	} else {
		result := output{
			TestType:      "test for differences",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   actual,
			ExpectedValue: "not nil",
		}
		OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	}
}
func TestCheckReturnType(actual interface{}, expected interface{}, request string) bool {
	if reflect.TypeOf(actual).String() == expected {
		result := output{
			TestType:      "test for matching return types",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   reflect.TypeOf(actual).String(),
			ExpectedValue: expected,
		}
		OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return true
	} else {
		result := output{
			TestType:      "test for matching return types",
			TestPosition:  0, // You might want to pass this as an argument or calculate it dynamically
			ActualValue:   reflect.TypeOf(actual).String(),
			ExpectedValue: expected,
		}
		OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\n", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue)
		fmt.Println(OutputString)
		return false
	}
}
