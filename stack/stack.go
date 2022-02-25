/*
Package stack provides a stack type and associated operations.

Copyright (c) 2012 Caleb Doxsey

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

//New creates a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

//Len returns the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

//Peek views the top item on the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

//Pop pops the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

//Push pushes a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
