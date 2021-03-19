package main

import "fmt"

type Instruction string

const (
	Right     Instruction = ">"
	Left      Instruction = "<"
	Increment Instruction = "+"
	Decrement Instruction = "-"
	Output    Instruction = "."
	Input     Instruction = ","
	LoopStart Instruction = "["
	LoopEnd   Instruction = "]"
)

var AllowedInstructions = []Instruction{Right, Left, Increment, Decrement, Output, Input, LoopStart, LoopEnd}

func IsValidInstruction(r rune) bool {
	x := Instruction(r)
	for _, ins := range AllowedInstructions {
		if x == ins {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("hello")
}
