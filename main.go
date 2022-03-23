package main

import (
	"os"

	"github.com/yelyzaveta-mykhalik/brainfuck/brainfuck"
)

func main() {

	code := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	compiler := brainfuck.NewCompiler(string(code))
	instructions := compiler.Compile()

	m := brainfuck.NewInterpreter(instructions, os.Stdin, os.Stdout)
	m.Execute()

}
