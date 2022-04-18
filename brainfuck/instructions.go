package brainfuck

import "fmt"

//Block of constants which represents aliases for chars of
//each command
const (
	Plus       byte = '+'
	Minus      byte = '-'
	RightShift byte = '>'
	LeftShift  byte = '<'
	WriteChar  byte = '.'
	ReadChar   byte = ','
	LoopStart  byte = '['
	LoopEnd    byte = ']'
)

//sizeOfCell represents capacity of cell, you can modify it,
//if you need
const sizeOfCell = 1000

//memoryCell struct represents structure and functionality
//of real cell in memory
type memoryCell struct {
	cell [sizeOfCell]byte
	ptr  int
}

//instruction interface implements compile()
//Via this method we can add new functionality to Brainfuck interpreter
type instruction interface {
	compile(m *memoryCell)
}

//plus type relates to char '+'
type plus struct{}

//increments memory on 1 point
func (pls plus) compile(m *memoryCell) {
	m.cell[m.ptr]++
}

//minus type relates to char '-'
type minus struct{}

//decrements memory on 1
func (mns minus) compile(m *memoryCell) {
	m.cell[m.ptr]--
}

//shiftingRight type relates to char '>'
type shiftingRight struct{}

//shift pointer to the right on 1
func (sr shiftingRight) compile(m *memoryCell) {
	m.ptr++
	m.ptr %= sizeOfCell
}

//shiftingLeft type relates to char '<'
type shiftingLeft struct{}

//shift pointer to the left on 1
func (sl shiftingLeft) compile(m *memoryCell) {
	m.ptr--
	if m.ptr < 0 {
		m.ptr = sizeOfCell - 1
	}
}

//writeChar type relates to char '.'
type writeChar struct{}

//output char into terminal
func (w writeChar) compile(m *memoryCell) {
	fmt.Printf("%c", m.cell[m.ptr])
}

//readChar type relates to char ','
type readChar struct{}

//read char from input in terminal
func (r readChar) compile(m *memoryCell) {
	fmt.Scanf("%c", &m.cell[m.ptr])
}

//startingLoop type relates to char '[' and contains slice
//to execute instructions in loop braces
type startingLoop struct {
	loopStack []instruction
}

//iterates through all instructions in loopStack
//and compile them
func (stl startingLoop) compile(m *memoryCell) {
	for m.cell[m.ptr] != 0 {
		for _, innerItem := range stl.loopStack {
			innerItem.compile(m)
		}
	}
}

//endingLoop type relates to char ']'
type endingLoop struct{}

//method for the end of the loop
func (endl endingLoop) compile(m *memoryCell) {}

//Instructions map contains the pairs of the instructions and
//the functions that implements them
var instructions = map[byte]instruction{
	Plus:       plus{},
	Minus:      minus{},
	RightShift: shiftingRight{},
	LeftShift:  shiftingLeft{},
	WriteChar:  writeChar{},
	ReadChar:   readChar{},
	LoopStart:  startingLoop{},
	LoopEnd:    endingLoop{},
}
