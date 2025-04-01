package main

import (
	"testing"
	"os"
	"bytes"
	"fmt"
)

func TestMainFunction(t *testing.T) {
	// Capture the output of the main function
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// Restore the original stdout
	w.Close()
	os.Stdout = rescueStdout

	out, _ := io.ReadAll(r)

	// Check if the output of the main function is as expected
	if string(out) != "hello world\n" {
		t.Errorf("Unexpected output. Expected 'hello world\n', got %s", out)
	}
}