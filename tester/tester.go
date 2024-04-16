// tester.go
package tester

import (
	"fmt"
	"reflect"
	"runtime"
)

var ExecutedTests []output
var RunTests = true    // export variables
var ShowOutput = false // export variables
var ShowDetailedPassed = false

type output struct {
	TestType      string
	ActualValue   interface{}
	ExpectedValue interface{}
	TestLine      int
	filePath      string
	status        bool
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

// Function to print ExecutedTests
func PrintExecutedTests() {
	fmt.Println("\nTests Executed:")
	var failedtests []output
	var passedCount, failedCount int
	for _, test := range ExecutedTests {
		if test.status {
			passedCount++
		} else {
			failedtests = append(failedtests, test)
			failedCount++
		}
	}
	fmt.Printf("Passed Tests: %d\n", passedCount)
	fmt.Printf("Failed Tests: %d\n", failedCount)
	for _, test := range failedtests {
		OutputString := fmt.Sprintf("\n%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorRed, ColorReset, test.TestType, test.ActualValue, test.ExpectedValue, test.TestLine, test.filePath)
		fmt.Println(OutputString)
	}
}

func TestEqual(actual interface{}, expected interface{}, request string) bool {
	if RunTests {
		_, filePath, line, _ := runtime.Caller(1)
		if actual == expected {
			result := output{
				TestType:      "test for Equality",
				ActualValue:   actual,
				ExpectedValue: expected,
				TestLine:      line,
				filePath:      filePath,
				status:        true,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				if ShowDetailedPassed {
					OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
					fmt.Println(OutputString)
				} else {
					OutputString := fmt.Sprintf("%sTest Passed%s", ColorGreen, ColorReset)
					fmt.Println(OutputString)
				}
			}
			return true
		} else {
			result := output{
				TestType:      "test for Equality",
				ActualValue:   actual,
				ExpectedValue: expected,
				TestLine:      line,
				filePath:      filePath,
				status:        false,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
				fmt.Println(OutputString)
			}
			return false
		}
	} else {
		return false
	}
}

func TestDifferent(actual interface{}, expected interface{}, request string) bool {
	if RunTests {
		_, filePath, line, _ := runtime.Caller(1)
		if actual != expected {
			result := output{
				TestType:      "test for difference",
				ActualValue:   actual,
				ExpectedValue: expected,
				TestLine:      line,
				filePath:      filePath,
				status:        true,
			}

			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				if ShowDetailedPassed {
					OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
					fmt.Println(OutputString)
				} else {
					OutputString := fmt.Sprintf("%sTest Passed%s", ColorGreen, ColorReset)
					fmt.Println(OutputString)
				}
			}
			return true
		} else {
			result := output{
				TestType:      "test for differences",
				ActualValue:   actual,
				ExpectedValue: expected,
				TestLine:      line,
				filePath:      filePath,
				status:        false,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
				fmt.Println(OutputString)
			}
			return false
		}
	} else {
		return false
	}
}
func TestNil(actual interface{}, request string) bool {
	if RunTests {
		_, filePath, line, _ := runtime.Caller(1)
		if actual == nil {
			result := output{
				TestType:      "test for Nil return",
				ActualValue:   actual,
				ExpectedValue: nil,
				TestLine:      line,
				filePath:      filePath,
				status:        true,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				if ShowDetailedPassed {
					OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
					fmt.Println(OutputString)
				} else {
					OutputString := fmt.Sprintf("%sTest Passed%s", ColorGreen, ColorReset)
					fmt.Println(OutputString)
				}
			}
			return true
		} else {
			result := output{
				TestType:      "test for nil return",
				ActualValue:   actual,
				ExpectedValue: nil,
				TestLine:      line,
				filePath:      filePath,
				status:        false,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
				fmt.Println(OutputString)
			}
			return true
		}
	} else {
		return false
	}
}
func TestNotNil(actual interface{}, request string) bool {
	if RunTests {
		_, filePath, line, _ := runtime.Caller(1)
		if actual != nil {
			result := output{
				TestType:      "test for not nill",
				ActualValue:   actual,
				ExpectedValue: "not nil",
				TestLine:      line,
				filePath:      filePath,
				status:        true,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				if ShowDetailedPassed {
					OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
					fmt.Println(OutputString)
				} else {
					OutputString := fmt.Sprintf("%sTest Passed%s", ColorGreen, ColorReset)
					fmt.Println(OutputString)
				}
			}
			return true
		} else {
			result := output{
				TestType:      "test for not nill",
				ActualValue:   actual,
				ExpectedValue: "not nil",
				TestLine:      line,
				filePath:      filePath,
				status:        false,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
				fmt.Println(OutputString)
			}
			return true
		}
	} else {
		return false
	}
}
func TestCheckReturnType(actual interface{}, expected interface{}, request string) bool {
	if RunTests {
		_, filePath, line, _ := runtime.Caller(1)
		if reflect.TypeOf(actual).String() == expected {
			result := output{
				TestType:      "test for matching return types",
				ActualValue:   reflect.TypeOf(actual).String(),
				ExpectedValue: expected,
				TestLine:      line,
				filePath:      filePath,
				status:        true,
			}
			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				if ShowDetailedPassed {
					OutputString := fmt.Sprintf("%sTest Passed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorGreen, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
					fmt.Println(OutputString)
				} else {
					OutputString := fmt.Sprintf("%sTest Passed%s", ColorGreen, ColorReset)
					fmt.Println(OutputString)
				}
			}
			return true
		} else {
			result := output{
				TestType:      "test for matching return types",
				ActualValue:   reflect.TypeOf(actual).String(),
				ExpectedValue: expected,
				TestLine:      line,
				filePath:      filePath,
				status:        false,
			}

			ExecutedTests = append(ExecutedTests, result)
			if ShowOutput {
				OutputString := fmt.Sprintf("%sTest Failed%s\nTest Type: %s\nActual: %v\nExpected: %v\nLine: %d::%s", ColorRed, ColorReset, result.TestType, result.ActualValue, result.ExpectedValue, result.TestLine, result.filePath)
				fmt.Println(OutputString)
			}
			return false
		}
	} else {
		return false
	}
}
