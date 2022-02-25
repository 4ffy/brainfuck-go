//Package tape provides a Tape type consisting of a tape of infinite cells
//that hold numerical values and a pointer to a certain cell for reading and
//writing.

package tape

import "fmt"

type Tape struct {
	maxval uint
	cell   uint
	cells  []uint
}

//NetTape initializes a new Tape.
func NewTape(width uint) *Tape {
	t := new(Tape)
	t.cells = make([]uint, 1)
	t.maxval = 1 << width

	return t
}

//MoveLeft moves one cell left.
func (t *Tape) MoveLeft() error {
	if t.cell == 0 {
		return fmt.Errorf("cannot move left: out of bounds")
	}
	t.cell--
	return nil
}

//MoveRight moves one cell right, resizing the tape if necessary.
func (t *Tape) MoveRight() error {
	if t.cell+1 == uint(len(t.cells)) {
		t.cells = append(t.cells, 0)
	}
	t.cell++
	return nil
}

//Increment adds one to the current cell, wrapping if limit reached.
func (t *Tape) Increment() {
	t.cells[t.cell] = (t.cells[t.cell] + 1) % t.maxval
}

//Increment subtracts one from current cell, wrapping if limit reached.
func (t *Tape) Decrement() {
	if t.cells[t.cell] > 0 {
		t.cells[t.cell]--
	} else {
		t.cells[t.cell] = t.maxval - 1
	}
}

//SetCell sets the current cell to an arbitrary value.
func (t *Tape) SetCell(value uint) error {
	if value >= t.maxval {
		return fmt.Errorf(
			"value %v is greater than max cell size %v",
			value, t.maxval-1,
		)
	}

	t.cells[t.cell] = value
	return nil
}

//GetCell returns the value of the current cell as an integer.
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
