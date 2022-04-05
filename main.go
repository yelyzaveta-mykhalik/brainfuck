package main

import (
	"github.com/yelyzaveta-mykhalik/brainfuck/brainfuck"
)

func main() {

	const helloWorld = "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	brainfuck.Execute(helloWorld)

}
