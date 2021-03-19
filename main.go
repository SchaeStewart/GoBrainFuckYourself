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

type loopSymbol struct {
	ins      Instruction
	position int
	depth    int
}

func FindLoops(input []Instruction) map[int]int {
	output := make(map[int]int)
	tmp := make([]loopSymbol, 0)
	depth := 0

	for i, ins := range input {
		if Instruction(ins) == LoopStart {
			depth++
			tmp = append(tmp, loopSymbol{
				ins:      ins,
				position: i,
				depth:    depth,
			})
		} else if Instruction(ins) == LoopEnd {
			tmp = append(tmp, loopSymbol{
				ins:      ins,
				position: i,
				depth:    depth,
			})
			depth--
		}
	}

	for i, val := range tmp {
		if val.ins == LoopStart {
			for _, x := range tmp[i+1:] {
				if x.depth == val.depth {
					output[val.position] = x.position
					output[x.position] = val.position
					break
				}
			}
		}
	}

	return output
}

func Run(input []Instruction) {
	stack := make([]int, 50)
	pointer := 0 
	loopPositions := FindLoops(input)

	for i := 0; i < len(input); i++ {
		ins := input[i]
		switch ins {
		case Right:
			pointer++
		case Left:
			pointer--
		case Increment:
			stack[pointer] += 1
		case Decrement:
			stack[pointer] -= 1
		case Output:
			fmt.Print(string(stack[pointer]))
		case Input:
			// TODO:
		case LoopStart:
			if stack[pointer] == 0 {
				i = loopPositions[i]
			}
		case LoopEnd:
			if stack[pointer] != 0 {
				i = loopPositions[i]
			}
		}
	}
}

func main() {
	fmt.Println("hello")
}
