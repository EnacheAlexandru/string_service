package main

import (
	"os"
	"testing"
)

/// AAA pattern (Arrange-Act-Assert)

func TestInputFileDoesNotExists(t *testing.T) {
	defer func() {
		recover()
	}()

	// Arrange
	finput := "input_test_wrong.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)

	// Assert
	t.Errorf("Input file exists. No panic")
}

func TestTransofrmation(t *testing.T) {
	// Arrange
	finput := "input_test.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)

	data, err := os.ReadFile(foutput)

	if err != nil {
		panic("Output file not found or failed to open!")
	}

	// =!!! because maps are unordered, each run generates a random order for the groups !!!=
	// the test input has only 2 groups, so there are only 2 possibilites
	expected1 := "A:\nAnita, anita@email.com, California\nAron, aron.bla@email.com, California\n\nC:\nCosmin, kox@bla.com, Giurgiu\n"
	expected2 := "C:\nCosmin, kox@bla.com, Giurgiu\n\nA:\nAnita, anita@email.com, California\nAron, aron.bla@email.com, California\n"

	// Assert
	if string(data) != expected1 && string(data) != expected2 {
		t.Errorf("Wrong output")
	}
}
