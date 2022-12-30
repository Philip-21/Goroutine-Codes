package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_UpdateMessage(t *testing.T) {

	wg.Add(1)
	go updateMessage("hello test")
	wg.Wait()

	if msg != "hello test" {
		t.Errorf("String error ")
	}
}

func Test_PrintMessage(t *testing.T) {

	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "hello test"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	out := string(result)
	os.Stdout = stdOut

	if !strings.Contains(out, "hello test") {
		t.Errorf("error in printing message")
	}
}

func Test_Main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	out := string(result)
	os.Stdout = stdOut

	if !strings.Contains(out, "Hello, Universe") {
		t.Errorf("String error ")
	}

	if !strings.Contains(out, "Hello, Cosmos") {
		t.Errorf("String error ")
	}
	if !strings.Contains(out, "Hello, World") {
		t.Errorf("String error ")
	}
}
