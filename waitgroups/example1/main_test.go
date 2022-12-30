package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout //program prints to os std

	r, w, _ := os.Pipe() //Pipe returns a connected pair of Files;
	os.Stdout = w        //defining the writer

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg) //always add pointer to waitgroups when defining them

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r) //r reads the result
	output := string(result)   ///cast the result into a string
	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected to find epsilon , but not there")
	}

}
