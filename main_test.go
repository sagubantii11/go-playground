// FILE: go-playground/main_test.go
package main

import (
	"bytes"
	"log"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Capture the output of the main function
	var buf bytes.Buffer
	log.SetOutput(&buf)
	main()
	log.SetOutput(nil)

	// Expected output
	expected := "Hello World!\n"

	// Compare the captured output with the expected output
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}
