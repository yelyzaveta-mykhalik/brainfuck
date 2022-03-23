package brainfuck

type Compiler struct {
	code       string
	codeLength int
	position   int

	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		codeLength:   len(code),
		instructions: []*Instruction{},
	}
}

func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{}

	for c.position < c.codeLength {
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

func (c *Compiler) CompileInstruction(char byte, insType InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.CountArgs(insType, count)
}

func (c *Compiler) CountArgs(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
