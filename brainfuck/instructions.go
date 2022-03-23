package brainfuck

//Alias on type byte for instrutctions
type InsType byte

const (
	Plus       InsType = '+'
	Minus      InsType = '-'
	RightShift InsType = '>'
	LeftShift  InsType = '<'
	WriteChar  InsType = '.'
	ReadChar   InsType = ','
	LoopStart  InsType = '['
	LoopEnd    InsType = ']'
)

type Instruction struct {
	Type     InsType
	Argument int
}
