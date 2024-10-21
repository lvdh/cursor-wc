// wc_test.go
package main

import (
	"os"
	"testing"
)

func TestCountWords(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write test content to the file
	content := "Hello World\nThis is a test file.\n"
	if _, err := tempFile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// Test the countWords function
	expected := 7
	actual, err := countWords(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
