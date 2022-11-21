package main

import (
	"os"
	"testing"
)

/// AAA pattern (Arrange-Act-Assert)

func TestTransofrmationInputFileDoesNotExists(t *testing.T) {
	defer func() {
		if err := recover(); err != "Input file not found or failed to open!" {
			// Assert
			t.Errorf("Expected panic: Input file not found or failed to open!")
		}
	}()

	// Arrange
	finput := "input_test_wrong.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)
}

func TestTransofrmationInputLowercaseName(t *testing.T) {
	defer func() {
		if err := recover(); err != "The first letter of each line (header excluded) should be an uppercase letter!" {
			// Assert
			t.Errorf("Expected panic: The first letter of each line (header excluded) should be an uppercase letter!")
		}
	}()

	// Arrange
	finput := "input_name_test.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)
}

func TestTransofrmationInputWrongHeader(t *testing.T) {
	defer func() {
		if err := recover(); err != "The header should have the first field \"full_name\". The fields are lowercase, separated with comma and space." {
			// Assert
			t.Errorf("Expected panic: The header should have the first field \"full_name\". The fields are lowercase, separated with comma and space.")
		}
	}()

	// Arrange
	finput := "input_header_test.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)
}

func TestTransofrmationInputFileIsEmpty(t *testing.T) {
	// Arrange
	finput := "input_empty_test.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)

	data, _ := os.ReadFile(foutput)

	// Assert
	expected := ""

	if string(data) != expected {
		t.Errorf("Expected an empty string")
	}
}

func TestTransofrmation(t *testing.T) {
	// Arrange
	finput := "input_test.csv"
	foutput := "output_test.csv"

	// Act
	Transformation(finput, foutput)

	data, _ := os.ReadFile(foutput)

	// =!!! because maps are unordered, each run generates a random order for the groups !!!=
	// the test input has only 2 groups, so there are only 2 possibilites

	// Assert
	expected1 := "A:\nAnita, anita@email.com, California\nAron, aron.bla@email.com, California\n\nC:\nCosmin, kox@bla.com, Giurgiu\n"
	expected2 := "C:\nCosmin, kox@bla.com, Giurgiu\n\nA:\nAnita, anita@email.com, California\nAron, aron.bla@email.com, California\n"

	if string(data) != expected1 && string(data) != expected2 {
		t.Errorf("Wrong output")
	}
}
