# Brainfuck Interpreter

A simple [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter
written in Go. It takes a naive approach without any fancy optimization.

Source code is given from a file via the `-i` option. The program detects
whether a Brainfuck program takes input and will only prompt for input if
necessary. Cell size can be set to any arbitrary size up to at least 32 bits
via the `b` option.
