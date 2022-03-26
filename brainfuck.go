package bf

import (
	"errors"
	"fmt"
)

type Instruction struct {
	operator Operator
	operand  uint16
}

func (i Instruction) String() string {
	return fmt.Sprintf("Instruction{%s *%d}", i.operator, i.operand)
}

type Operator uint16

const (
	Noop Operator = iota
	Right
	Left
	Increment
	Decrement
	Print
	Read
	BeginLoop
	EndLoop
)

func (op Operator) String() string {
	switch op {
	case Noop:
		return "Noop"
	case Right:
		return "Right"
	case Left:
		return "Left"
	case Increment:
		return "Increment"
	case Decrement:
		return "Decrement"
	case Print:
		return "Print"
	case Read:
		return "Read"
	case BeginLoop:
		return "BeginLoop"
	case EndLoop:
		return "EndLoop"
	default:
		return "unknown"
	}
}

const dataSize = 65535

type BrainfuckCompiler struct {
	instructions []Instruction
	memory       []uint8
	output       []byte
	inputs       int
}

func NewCompiler() *BrainfuckCompiler {
	return &BrainfuckCompiler{memory: make([]uint8, dataSize)}
}

func (bf *BrainfuckCompiler) Compile(program string) error {
	var pc, jumpPc uint16 = 0, 0
	jumpStack := make([]uint16, 0)

	for _, c := range program {
		switch c {
		case '>':
			bf.instructions = append(bf.instructions, Instruction{Right, 0})
		case '<':
			bf.instructions = append(bf.instructions, Instruction{Left, 0})
		case '+':
			bf.instructions = append(bf.instructions, Instruction{Increment, 0})
		case '-':
			bf.instructions = append(bf.instructions, Instruction{Decrement, 0})
		case '.':
			bf.instructions = append(bf.instructions, Instruction{Print, 0})
		case ',':
			bf.instructions = append(bf.instructions, Instruction{Read, 0})
			bf.inputs++
		case '[':
			bf.instructions = append(bf.instructions, Instruction{BeginLoop, 0})
			jumpStack = append(jumpStack, pc)
		case ']':
			if len(jumpStack) == 0 {
				return errors.New("Compilation error")
			}
			jumpPc = jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]
			bf.instructions = append(bf.instructions, Instruction{EndLoop, jumpPc})
			bf.instructions[jumpPc].operand = pc
		default:
			pc--
		}
		pc++
	}
	if len(jumpStack) != 0 {
		return errors.New("Compilation error")
	}
	return nil
}

func (bf *BrainfuckCompiler) Execute(input []byte) error {
	pointer := uint16(0)
	readCount := 0

	if len(input) != bf.inputs {
		return errors.New("Not enough input characters")
	}

	for pc := 0; pc < len(bf.instructions); pc++ {
		switch bf.instructions[pc].operator {
		case Right:
			pointer++
		case Left:
			pointer--
		case Increment:
			bf.memory[pointer]++
		case Decrement:
			bf.memory[pointer]--
		case Print:
			bf.output = append(bf.output, byte(bf.memory[pointer]))
		case Read:
			bf.memory[pointer] = input[readCount]
			readCount++
		case BeginLoop:
			if bf.memory[pointer] == 0 {
				pc = int(bf.instructions[pc].operand)
			}
		case EndLoop:
			if bf.memory[pointer] > 0 {
				pc = int(bf.instructions[pc].operand)
			}
		case Noop:
		default:
			return errors.New("Invalid Operator")
		}
	}

	return nil
}

func (bf *BrainfuckCompiler) Output() string {
	return string(bf.output[:])
}

func (bf *BrainfuckCompiler) Clear() error {
	bf.memory = make([]uint8, dataSize)
	bf.instructions = []Instruction{}
	bf.output = []byte{}
	return nil
}
