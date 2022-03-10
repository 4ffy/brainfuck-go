/*
Package interpreter takes a Tape object and builds a brainfuck interpreter
using the tape as memory.

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

package interpreter

import (
	"brainfuck/stack"
	"brainfuck/tape"
	"fmt"
	"regexp"
	"strings"
)

//BFInterpreter provides a interpreter container type.
type BFInterpreter struct {
	memory *tape.Tape //Cell memory.
}

//token provides a container storing a BF instruction and how many times it is
//to be executed.
type token struct {
	op  byte //brainfuck instruction
	num uint //number of times to run
}

//New initializes a new brainfuck interpreter with cell width
//given in bits.
func New(width uint) *BFInterpreter {
	bf := new(BFInterpreter)
	bf.memory = tape.New(width)
	return bf
}

//Execute runs a brainfuck program from source and input strings, printing to
//stdout.
func (bf *BFInterpreter) Execute(source, input string) error {
	//Prepare for execution.
	bf.Reset()
	source = cleanSource(source)
	tokens := tokenize(source)
	loops, err := getLoops(tokens)
	if err != nil {
		return fmt.Errorf("creating loop map: %v", err)
	}

	//Loop through tokens.
	inpPtr := 0
	tokPtr := 0
	for tokPtr < len(tokens) {
		switch tokens[tokPtr].op {
		case '+':
			bf.memory.Add(tokens[tokPtr].num)
		case '-':
			bf.memory.Subtract(tokens[tokPtr].num)
		case '>':
			err = bf.memory.MoveRight(tokens[tokPtr].num)
			if err != nil {
				return fmt.Errorf("pos %v op >: %v", tokPtr, err)
			}
		case '<':
			err = bf.memory.MoveLeft(tokens[tokPtr].num)
			if err != nil {
				return fmt.Errorf("pos %v op <: %v", tokPtr, err)
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
				tokPtr = loops[tokPtr]
			}
		case ']':
			if bf.memory.GetCell() != 0 {
				tokPtr = loops[tokPtr]
			}
		}
		tokPtr++
	}

	fmt.Print("\n")
	return nil
}

//Reset clears the memory tape, setting all cells to 0 and moving back to the
//first cell.
func (bf *BFInterpreter) Reset() {
	bf.memory.Reset()
}

//PrintDebug dumps the contents of the memory tape to stdout.
func (bf *BFInterpreter) PrintDebug() {
	bf.memory.PrintDebug()
}

//cleanSource removes characters from a Brainfuck source that serve no purpose,
//such as comments and newlines.
func cleanSource(source string) string {
	re := regexp.MustCompile(`[^+-<>.,\[\]]`) //Remove non +=<>[],.
	return re.ReplaceAllString(source, "")
}

//tokenize returns a list of tokens through run-length encoding a source
//string. Only +-<> instructions get properly RLE'd since the other operations
//provide little benefit.
func tokenize(source string) []token {
	tokens := make([]token, 0)
	count := uint(1)

	//Loop through source.
	for i := range source[:len(source)-1] {
		if source[i] == source[i+1] &&
			strings.ContainsRune("+-<>", rune(source[i])) {
			count++
		} else {
			tokens = append(tokens, token{source[i], count})
			count = 1
		}
	}

	tokens = append(tokens, token{source[len(source)-1], count}) //Final instruction
	return tokens
}

//getLoops reads tokenized source, returning a map of its loops.
func getLoops(tokens []token) (map[int]int, error) {
	loops := make(map[int]int)
	lStack := stack.New()

	//Loop through source.
	for ptr, tok := range tokens {
		switch tok.op {
		case '[':
			lStack.Push(ptr)
		case ']':
			if lStack.Len() == 0 {
				return nil, fmt.Errorf(
					"pos %v op ]: close loop without matching open", ptr)
			}

			//Add loops to map.
			loops[ptr] = lStack.Pop().(int)
			loops[loops[ptr]] = ptr
		}
	}

	//The stack should be empty at this point.
	if lStack.Len() > 0 {
		return nil, fmt.Errorf(
			"pos %v op [: open loop without matching close", lStack.Pop())
	}

	return loops, nil
}
