/*
Package tape provides a Tape type consisting of a tape of infinite cells
that hold numerical values and a pointer to a certain cell for reading and
writing.

Copyright (c) 2022 Cameron Norton

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package tape

import "fmt"

//Tape provides a memory tape consisting of a list of cells and a pointer to an
//active cell.
type Tape struct {
	maxval uint   //maximum cell value
	cell   uint   //current cell
	cells  []uint //cell memory
}

//New initializes a new Tape.
func New(width uint) *Tape {
	t := new(Tape)
	t.cells = make([]uint, 1)
	t.maxval = 1 << width
	return t
}

//MoveLeft moves the active cell n cells to the left.
func (t *Tape) MoveLeft(n uint) error {
	if n > t.cell {
		return fmt.Errorf("move left %v cells: out of bounds", n)
	}
	t.cell -= n
	return nil
}

//MoveRight moves the active cell n cells right, resizing the tape if
//necessary.
func (t *Tape) MoveRight(n uint) error {
	if t.cell+n >= uint(len(t.cells)) {
		extend := 1 + t.cell + n - uint(len(t.cells))
		t.cells = append(t.cells, make([]uint, extend)...)
	}
	t.cell += n
	return nil
}

//Add adds n to the active cell, wrapping if limit reached.
func (t *Tape) Add(n uint) {
	t.cells[t.cell] = (t.cells[t.cell] + n) % t.maxval
}

//Subtract subtracts n from the active cell, wrapping if limit reached.
func (t *Tape) Subtract(n uint) {
	if n > t.cells[t.cell] {
		t.cells[t.cell] = t.maxval - (n - t.cells[t.cell])
	} else {
		t.cells[t.cell] -= n
	}
}

//SetCell sets the active cell to an arbitrary value.
func (t *Tape) SetCell(value uint) error {
	if value >= t.maxval {
		return fmt.Errorf(
			"value %v is greater than max cell size %v",
			value, t.maxval-1)
	}

	t.cells[t.cell] = value
	return nil
}

//GetCell returns the value of the active cell as an integer.
func (t *Tape) GetCell() uint {
	return t.cells[t.cell]
}

//Reset clears the tape, setting all cells to 0 and moving back to the first
//cell.
func (t *Tape) Reset() {
	for i := range t.cells {
		t.cells[i] = 0
	}
	t.cell = 0
}

//PrintDebug dumps the contents of the tape to stdout.
func (t *Tape) PrintDebug() {
	for i := range t.cells {
		fmt.Printf("%v\t", t.cells[i])
	}
	fmt.Printf("\n")
}
