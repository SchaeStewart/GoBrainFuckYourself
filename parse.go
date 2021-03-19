package main

func Parse(input string) []Instruction {
	instructions := make([]Instruction, 0, len(input))
	for _, r := range input {
		if IsValidInstruction(r) {
			instructions = append(instructions, Instruction(r))
		}
	}
	return instructions
}
