package brainfuck

import (
	"bytes"
	"testing"
)

func TestCompile(t *testing.T) {
	input := `
	+++++++
	---
	+++++++
	>>>
	<<<<<<<
	`
	expected := []*Instruction{
		&Instruction{Plus, 7},
		&Instruction{Minus, 3},
		&Instruction{Plus, 7},
		&Instruction{RightShift, 3},
		&Instruction{LeftShift, 7},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("Wrong length of bytecode. Expected: %+v. Given: %+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("Wrong operation. Expected: %+v, Given: %+v", op, bytecode[i])
		}
	}
}

func TestCompileLoops(t *testing.T) {
	input := `+[+[+]+]+`
	expected := []*Instruction{
		&Instruction{Plus, 1},
		&Instruction{LoopStart, 7},
		&Instruction{Plus, 1},
		&Instruction{LoopStart, 5},
		&Instruction{Plus, 1},
		&Instruction{LoopEnd, 3},
		&Instruction{Plus, 1},
		&Instruction{LoopEnd, 1},
		&Instruction{Plus, 1},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("Wrong length of bytecode. Expected: %+v. Given: %+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("Wrong length of bytecode. Expected: %+v. Given: %+v", op, bytecode[i])
		}
	}
}

func TestCompileAllFunctionality(t *testing.T) {
	input := `+++[---[+]>>>]<<<`
	expected := []*Instruction{
		&Instruction{Plus, 3},
		&Instruction{LoopStart, 7},
		&Instruction{Minus, 3},
		&Instruction{LoopStart, 5},
		&Instruction{Plus, 1},
		&Instruction{LoopEnd, 3},
		&Instruction{RightShift, 3},
		&Instruction{LoopEnd, 1},
		&Instruction{LeftShift, 3},
	}

	compiler := NewCompiler(input)
	bytecode := compiler.Compile()

	if len(bytecode) != len(expected) {
		t.Fatalf("Wrong length of bytecode. Expected: %+v. Given: %+v",
			len(expected), len(bytecode))
	}

	for i, op := range expected {
		if *bytecode[i] != *op {
			t.Errorf("Wrong length of bytecode. Expected: %+v. Given: %+v", op, bytecode[i])
		}
	}
}

func TestIncrement(t *testing.T) {
	compiler := NewCompiler("+++++")
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("A cell wasn't correctly increment. Given: %d", m.memory[0])
	}
}

func TestDecrement(t *testing.T) {
	input := "++++++++++-----"
	compiler := NewCompiler(input)
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	if m.memory[0] != 5 {
		t.Errorf("A cell wasn't correctly decrement. Given: %d", m.memory[0])
	}
}

func TestIncrementingDataPointer(t *testing.T) {
	compiler := NewCompiler("+>++>+++")
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	for i, expected := range []int{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("A memory[%d] has incorrect value. Expected: %d. Given: %d",
				i, expected, m.memory[0])
		}
	}
}

func TestDecrementDataPointer(t *testing.T) {
	compiler := NewCompiler(">>+++<++<+")
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, new(bytes.Buffer), new(bytes.Buffer))

	m.Execute()

	for i, expected := range []int{1, 2, 3} {
		if m.memory[i] != expected {
			t.Errorf("A memory[%d] has incorrect value. Expected: %d. Given: %d",
				i, expected, m.memory[0])
		}
	}
}

func TestReadChar(t *testing.T) {
	in := bytes.NewBufferString("ABCDEF")
	out := new(bytes.Buffer)

	compiler := NewCompiler(",>,>,>,>,>,>")
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, in, out)

	m.Execute()

	expectedMemory := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, expected := range expectedMemory {
		if m.memory[i] != expected {
			t.Errorf("A memory[%d] has incorrect value. Expected: %d. Given: %d",
				i, expected, m.memory[0])
		}
	}
}

func TestPutChar(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	compiler := NewCompiler(".>.>.>.>.>.>")
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, in, out)

	setupMemory := []int{
		int('A'),
		int('B'),
		int('C'),
		int('D'),
		int('E'),
		int('F'),
	}

	for i, value := range setupMemory {
		m.memory[i] = value
	}

	m.Execute()

	output := out.String()
	if output != "ABCDEF" {
		t.Errorf("Incorrect output. Given:", output)
	}

}

const HelloWorld = `++++++++[>++++[>++>+++>+++>+<<<<-]>+> +>->>+[<]<-]>>.>---.+++++++ ..+ ++.>>.<-.<.+++.------.--------.>>+.>++.`

func TestHelloWorld(t *testing.T) {
	in := bytes.NewBufferString("")
	out := new(bytes.Buffer)

	compiler := NewCompiler(HelloWorld)
	instructions := compiler.Compile()

	m := NewInterpreter(instructions, in, out)

	m.Execute()

	output := out.String()
	if output != "Hello World!\n" {
		t.Errorf("Incorrect output. Given: %q", output)
	}
}
