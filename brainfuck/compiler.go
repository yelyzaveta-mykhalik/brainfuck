package brainfuck

//Compiler represents structure of compiler
//Where code is brainfuck program
//Position is position of character in brainfuck program
//Instructions - slice that refer to Instruction struct
type Compiler struct {
	code     string
	position int

	instructions []*Instruction
}

//NewCompiler creates instanse of Compiler struct
func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		instructions: []*Instruction{},
	}
}

//Compile iteraits through each symbol of brainfuck program and
//execute commands from the program,
//depending on the character of which they are represented
func (c *Compiler) Compile() []*Instruction {
	var loopStack = []int{}

	for c.position < len(c.code) {
		current := c.code[c.position]

		switch current {
		case '[':
			insPos := c.CountArgs(LoopStart, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			closeInstructionPos := c.CountArgs(LoopEnd, openInstruction)
			// Patch the old LoopStart ("[") instruction with new position
			c.instructions[openInstruction].Argument = closeInstructionPos
		case '+':
			c.CompileInstruction('+', Plus)
		case '-':
			c.CompileInstruction('-', Minus)
		case '<':
			c.CompileInstruction('<', LeftShift)
		case '>':
			c.CompileInstruction('>', RightShift)
		case '.':
			c.CompileInstruction('.', WriteChar)
		case ',':
			c.CompileInstruction(',', ReadChar)
		}

		c.position++
	}

	return c.instructions
}

//
func (c *Compiler) CompileInstruction(char byte, insType InsType) {
	count := 1

	for c.position < len(c.code)-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.CountArgs(insType, count)
}

//CountArgs return amount of symbols in brainfuck program
func (c *Compiler) CountArgs(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
