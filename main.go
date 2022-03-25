package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/yelyzaveta-mykhalik/brainfuck/brainfuck"
)

func main() {

	//Getting file name from CLI, reading that file
	//and executing context of file
	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}

	compiler := brainfuck.NewCompiler(string(code))
	instructions := compiler.Compile()

	m := brainfuck.NewInterpreter(instructions, os.Stdin, os.Stdout)
	m.Execute()

}
