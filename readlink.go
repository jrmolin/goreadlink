// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {

    // `os.Args` provides access to raw command-line
    // arguments. Note that the first value in this slice
    // is the path to the program, and `os.Args[1:]`
    // holds the arguments to the program.
	argsWithoutProg := os.Args[1:]

	for _, value := range argsWithoutProg {
		abs, _ := filepath.Abs(value)
		fmt.Println(abs)
	}

}

