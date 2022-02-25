# Brainfuck Interpreter

A simple [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter
written in Go. It takes a naive approach without any fancy optimization.

Source code is given from a file via the `-i` option. The program detects
whether a Brainfuck program takes input and will only prompt for input if
necessary. Cell size can be set to any arbitrary size up to at least 32 bits
via the `-b` option.

## Caveats
Every Brainfuck interpreter has its idiosyncrasies. Here's some of the
relevant ones for this program.
- Cells wrap around when they go below zero or above the max cell size.
- Cells are set to 0 when a `,` operation is read and there is no input left.
- Input strings are given on a single line and terminated by a newline, which
  is not included when passed to the interpreter.
