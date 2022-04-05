package brainfuck

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCompile(t *testing.T) {
	m := memoryCell{}

	testSet := []struct {
		ins             instruction
		ptr             int
		value, expected byte
	}{
		{new(plus), 0, 1, 2},
		{new(plus), 0, 100, 15},
		{new(minus), 7, 0, 250},
		{new(minus), 1, 5, 3},
		{new(shiftingRight), 0, 2, 4},
		{new(shiftingRight), 199, 10, 0},
		{new(shiftingLeft), 0, 0, 89},
		{new(shiftingLeft), 6, 0, 4},
	}

	for _, test := range testSet {
		m.ptr = test.ptr
		m.cell[m.ptr] = test.value

		test.ins.compile(&m)

		if m.cell[m.ptr] != test.expected {
			t.Errorf("Expected: %d.Given:  %d", test.expected, m.cell[m.ptr])
		}
	}
}

func TestWriteChar(t *testing.T) {
	var ins writeChar
	var m memoryCell

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	m.ptr = 0
	m.cell[m.ptr] = 33

	ins.compile(&m)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	if out[0] != 33 {
		t.Errorf("Expected: %d. Given: %d", 100, m.cell[m.ptr])
	}
}

func TestLoop(t *testing.T) {
	m := memoryCell{}

	testSet := []struct {
		ins             instruction
		ptr             int
		value, expected byte
	}{
		{
			startingLoop{
				[]instruction{plus{}, minus{}, shiftingLeft{}, shiftingRight{}},
			},
			0,
			1,
			2},
	}

	for _, test := range testSet {
		m.ptr = test.ptr
		m.cell[m.ptr] = test.value

		test.ins.compile(&m)

		if m.cell[m.ptr] != test.expected {
			t.Errorf("Expected: %d. Given: %d", test.expected, m.cell[m.ptr])
		}
	}
}

const testString = "++++++++++[>+>+++>+++++++>++++++++++<<<<-]>>+++++++++++++++++++.-.."

func TestInterpret(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Execute(testString)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOut

	if out[0] != 100 {
		t.Errorf("Expected: 100. Equals: %d", out[0])
	}
}
