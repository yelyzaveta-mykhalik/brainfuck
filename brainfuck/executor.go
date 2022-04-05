package brainfuck

//Execute is the main function of program that takes as an argument
//string in Brainfuck and execute every instruction from this string
func Execute(programCode string) {

	//currentStack contains slice for executing commands in loop
	var currentStack = []startingLoop{}

	//In this loop, depending on type of symbol from string, we push instruction
	//in currentStack
	for i := 0; i < len(programCode); i++ {
		var currentInstruction = instructions[programCode[i]]

		switch t := currentInstruction.(type) {
		case startingLoop:
			//appending loop into currentStack
			currentStack = append([]startingLoop{t}, currentStack...)
		case endingLoop:
			currentStack[1].loopStack = append(currentStack[1].loopStack, currentStack[0])
			//popping elements of currentStack
			currentStack = currentStack[1:(len(currentStack) + 1)]
			//appending instructions to currentStack but into the loopStack
		case plus, minus, shiftingRight, shiftingLeft, writeChar, readChar:
			currentStack[0].loopStack = append(currentStack[0].loopStack, currentInstruction)
		}
	}

	//executing final slice of instructions
	interpret(currentStack[0].loopStack)
}

//interpret execute all instructions
func interpret(setOfInstructions []instruction) {
	var memory memoryCell

	for _, i := range setOfInstructions {
		i.compile(&memory)
	}
}
