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
//Via this method we can addd new functionality to Brainfuck interpreter
type instruction interface {
	compile(m *memoryCell)
}

//plus type relates to char '+'
type plus struct{}

//implements compile for type plus
func (pls plus) compile(m *memoryCell) {
	m.cell[m.ptr]++
}

//minus type relates to char '-'
type minus struct{}

//implements compile for type minus
func (mns minus) compile(m *memoryCell) {
	m.cell[m.ptr]--
}

//shiftingRight type relates to char '>'
type shiftingRight struct{}

//implements compile for type shiftRight
func (sr shiftingRight) compile(m *memoryCell) {
	m.ptr++
}

//shiftingLeft type relates to char '<'
type shiftingLeft struct{}

//implements compile for type shiftLeft
func (sl shiftingLeft) compile(m *memoryCell) {
	m.ptr--
}

//writeChar type relates to char '.'
type writeChar struct{}

//implements compile for type wryteChar
func (w writeChar) compile(m *memoryCell) {
	fmt.Printf("%c", m.cell[m.ptr])
}

//readChar type relates to char ','
type readChar struct{}

//implements compile for type readChar
func (r readChar) compile(m *memoryCell) {
	fmt.Scanf("%c", m.cell[m.ptr])
}

//startingLoop type relates to char '[' and contains slice
//to execute instructions in loop braces
type startingLoop struct {
	loopStack []instruction
}

//implements compile for type startingLoop, iterate through
//all instructions in loopStack and compile them
func (stl startingLoop) compile(m *memoryCell) {
	if m.cell[m.ptr] != 0 {
		for _, loopItem := range stl.loopStack {
			loopItem.compile(m)
		}
	}
}

//endingLoop type relates to char ']'
type endingLoop struct{}

//implements compile for type readChar
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
