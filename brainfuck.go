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

func hasInput(source string) bool {
	for _, v := range source {
		if v == ',' {
			return true
		}
	}
	return false
}

func main() {
	file := flag.String("i", "", "input source file")
	width := flag.Uint("b", 8, "size of cells in bits")
	flag.Parse()

	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatalf("could not read file %v: %v", file, err)
	}
	source := string(data)

	input := ""
	if hasInput(source) {
		in := bufio.NewReader(os.Stdin)
		input, err = in.ReadString('\n')
		if err != nil {
			log.Fatalf("could not read input: %v", err)
		}
	}

	bf := interpreter.NewBFInterpreter(*width)
	err = bf.Execute(source, input)
	if err != nil {
		log.Fatalf("execute brainfuck: %v", err)
	}
	fmt.Print("\n")
}
