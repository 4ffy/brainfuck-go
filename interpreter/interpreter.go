/*
Package interpreter takes a Tape object and builds a brainfuck interpreter
using the tape as memory.

Copyright (c) 2022 4ffy

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

package interpreter

import (
	"brainfuck/stack"
	"brainfuck/tape"
	"fmt"
	"regexp"
)

type BFInterpreter struct {
	memory *tape.Tape
}

//NewBFInterpreter initializes a new brainfuck interpreter with cell width
//given in bits.
func NewBFInterpreter(width uint) *BFInterpreter {
	bf := new(BFInterpreter)
	bf.memory = tape.NewTape(width)
	return bf
}

//getLoops reads brainfuck source, returning a map of its loops.
func getLoops(source string) (map[int]int, error) {
	loops := make(map[int]int)
	lStack := stack.New()

	for ptr, op := range source {
		switch op {
		case '[':
			lStack.Push(ptr)
		case ']':
			if lStack.Len() == 0 {
				return nil, fmt.Errorf(
					"pos %v op ]: close loop without matching open", ptr)
			}
			loops[ptr] = lStack.Pop().(int)
			loops[loops[ptr]] = ptr
		}
	}

	if lStack.Len() > 0 {
		return nil, fmt.Errorf(
			"pos %v op [: open loop without matching close", lStack.Pop())
	}

	return loops, nil
}

//Execute runs a brainfuck program from source and input strings, printing to
//stdout.
func (bf *BFInterpreter) Execute(source, input string) error {
	source = cleanSource(source)
	bf.Reset()

	loops, err := getLoops(source)
	if err != nil {
		return fmt.Errorf("creating loop map: %v", err)
	}

	var inpPtr int = 0
	var srcPtr int = 0
	for srcPtr < len(source) {
		switch source[srcPtr] {
		case '+':
			bf.memory.Increment()
		case '-':
			bf.memory.Decrement()
		case '>':
			err = bf.memory.MoveRight()
			if err != nil {
				return fmt.Errorf("pos %v op >: cannot move right: %v", srcPtr, err)
			}
		case '<':
			err = bf.memory.MoveLeft()
			if err != nil {
				return fmt.Errorf("pos %v op <: cannot move left: %v", srcPtr, err)
			}
		case '.':
			fmt.Printf("%c", rune(bf.memory.GetCell()))
		case ',':
			if inpPtr < len(input) {
				bf.memory.SetCell(uint(input[inpPtr]))
				inpPtr++
			}
		case '[':
			if bf.memory.GetCell() == 0 {
				srcPtr = loops[srcPtr]
			}
		case ']':
			if bf.memory.GetCell() != 0 {
				srcPtr = loops[srcPtr]
			}
		}
		srcPtr++
	}

	fmt.Print("\n")
	return nil
}

//Reset clears the memory tape, setting all cells to 0 and moving back to the
//first cell.
func (bf *BFInterpreter) Reset() {
	bf.memory.Reset()
}

//PrintDebug dumps the contents of the tape to stdout.
func (bf *BFInterpreter) PrintDebug() {
	bf.memory.PrintDebug()
}

func cleanSource(source string) string {
	re := regexp.MustCompile(`[^+-<>.,\[\]]`)
	return re.ReplaceAllString(source, "")
}
