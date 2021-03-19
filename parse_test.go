package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	// input := ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+."

	input := "><+-.,[]"

	output := Parse(input)

	assert.Equal(t, output, []Instruction{Right, Left, Increment, Decrement, Output, Input, LoopStart, LoopEnd})

}
