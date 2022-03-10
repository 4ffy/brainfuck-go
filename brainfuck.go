/*
Package main provides a main method for using the interpreter.

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

package main

import (
	"brainfuck/interpreter"
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//hasInput determines whether a Brainfuck source has any , instructions.
func hasInput(source string) bool {
	for _, v := range source {
		if v == ',' {
			return true
		}
	}
	return false
}

func main() {
	//Parse options.
	file := flag.String("i", "", "input source file")
	width := flag.Uint("b", 8, "size of cells in bits")
	flag.Parse()

	//Read source.
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("could not read file %v: %v", file, err)
	}
	source := string(data)

	//Read input, if necessary.
	input := ""
	if hasInput(source) {
		in := bufio.NewReader(os.Stdin)
		input, err = in.ReadString('\n')
		if err != nil {
			log.Fatalf("could not read input: %v", err)
		}
		input = input[:len(input)-1]
	}

	//Execute Brainfuck.
	bf := interpreter.New(*width)
	err = bf.Execute(source, input)
	if err != nil {
		log.Fatalf("execute brainfuck: %v", err)
	}
	fmt.Print("\n")
}
