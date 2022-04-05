package brainfuck

import "io"

type Interpreter struct {
	code []*Instruction
	ip   int

	memory  [30000]int
	dataPtr int

	input  io.Reader
	output io.Writer

	readBuf []byte
}

func NewInterpreter(instructions []*Instruction, in io.Reader, out io.Writer) *Interpreter {
	return &Interpreter{
		code:    instructions,
		input:   in,
		output:  out,
		readBuf: make([]byte, 1),
	}
}

func (intr *Interpreter) Execute() {
	for intr.ip < len(intr.code) {
		ins := intr.code[intr.ip]

		switch ins.Type {
		case Plus:
			intr.memory[intr.dataPtr] += ins.Argument
		case Minus:
			intr.memory[intr.dataPtr] -= ins.Argument
		case RightShift:
			intr.dataPtr += ins.Argument
		case LeftShift:
			intr.dataPtr -= ins.Argument
		case WriteChar:
			for i := 0; i < ins.Argument; i++ {
				intr.writeChar()
			}
		case ReadChar:
			for i := 0; i < ins.Argument; i++ {
				intr.readChar()
			}
		case LoopStart:
			if intr.memory[intr.dataPtr] == 0 {
				intr.ip = ins.Argument
				continue
			}
		case LoopEnd:
			if intr.memory[intr.dataPtr] != 0 {
				intr.ip = ins.Argument
				continue
			}
		}

		intr.ip++
	}
}

func (intr *Interpreter) readChar() {
	n, err := intr.input.Read(intr.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("Wrong num bytes had been read")
	}

	intr.memory[intr.dataPtr] = int(intr.readBuf[0])
}

func (intr *Interpreter) writeChar() {
	intr.readBuf[0] = byte(intr.memory[intr.dataPtr])

	n, err := intr.output.Write(intr.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("Wrong num bytes had been written")
	}
}
