package main

import (
	"tester/tester"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestEqual(t *testing.T) {
	result := tester.TestEqual(1, 1, "has to be equal")
	assert.Equal(t, result, true, "has to be true")
}
func TestTestDifferent(t *testing.T) {
	result := tester.TestDifferent(1, 2, "has to be different")
	assert.Equal(t, result, true, "has to be true")
}
func TestTestNil(t *testing.T) {
	result := tester.TestNil(nil, "has to be nil")
	assert.Equal(t, result, true, "has to be true")
}
func TestTestNotNil(t *testing.T) {
	result := tester.TestNotNil(1, "has not to be nil")
	assert.Equal(t, result, true, "has to be true")
}
func TestTestCheckReturnType(t *testing.T) {
	result := tester.TestCheckReturnType(1, "int", "has to be int")
	assert.Equal(t, result, true, "has to be true")
}
func TestForFailedTestEqual(t *testing.T) {
	// Failed test case: 1 is not equal to 2
	result := tester.TestEqual(1, 2, "has to be equal")
	assert.Equal(t, result, false, "has to be false")
}

func TestForFailedTestDifferent(t *testing.T) {
	// Failed test case: 1 is equal to 1
	result := tester.TestDifferent(1, 1, "has to be different")
	assert.Equal(t, result, false, "has to be false")
}

func TestForFailedTestNil(t *testing.T) {
	// Failed test case: "not nil" is not nil
	result := tester.TestNil("not nil", "has to be nil")
	assert.Equal(t, result, true, "has to be false")
}

func TestForFailedTestNotNil(t *testing.T) {
	// Failed test case: 1 is nil
	result := tester.TestNotNil(nil, "has not to be nil")
	assert.Equal(t, result, true, "has to be false")
}

func TestForFailedTestCheckReturnType(t *testing.T) {
	// Failed test case: 1 is not of type "string"
	result := tester.TestCheckReturnType(1, "string", "has to be int")
	assert.Equal(t, result, false, "has to be false")
}
