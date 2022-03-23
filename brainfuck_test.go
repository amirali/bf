package bf

import (
	"reflect"
	"testing"
)

func TestCompileOperators(t *testing.T) {
	compiler := NewCompiler()
	operators := []Operator{Right, Left, Increment, Decrement, Print, Read, BeginLoop, EndLoop}
	program := "><+-.,[]"
	compiler.Compile(program)
	for i := 0; i < len(program); i++ {
		if compiler.instructions[i].operator != operators[i] {
			t.Errorf("Expected Operator(%v) but got Operator(%v)", operators[i],
				compiler.instructions[i].operator)
		}
	}
}

func TestCompileLoopOperand(t *testing.T) {
	compiler := NewCompiler()
	program := "++++[--]"
	compiler.Compile(program)
	if compiler.instructions[5].operand != 0 {
		t.Errorf("Loop begin point operand must be %d but it's %d", 0, compiler.instructions[5].operand)
	}
	if compiler.instructions[7].operand != 4 {
		t.Errorf("End begin point operand must be %d but it's %d", 4, compiler.instructions[7].operand)
	}
}

func TestHelloWorld(t *testing.T) {
	compiler := NewCompiler()
	program := ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+."
	compiler.Compile(program)
	input := make([]byte, 0)
	compiler.Execute(input)
	if compiler.Output() != "Hello, World!" {
		t.Errorf("Code output must be %q but it's %q", "Hello, World!", compiler.Output())
	}
}

func TestWithInput(t *testing.T) {
	compiler := NewCompiler()
	program := ",."
	compiler.Compile(program)
	input := []byte{'c'}
	compiler.Execute(input)
	if compiler.Output() != "c" {
		t.Errorf("Code output must be %q but it's %q", "c", compiler.Output())
	}
}

func TestClearCompiler(t *testing.T) {
	compiler := NewCompiler()
	program := ">>[-]<<[->>+<<]"
	compiler.Compile(program)
	compiler.Execute([]byte{})
	err := compiler.Clear()
	if err != nil {
		t.Error("Error in clearing compiler")
	}
	freshMemory := make([]uint8, dataSize)
	if len(compiler.instructions) != 0 || compiler.Output() != "" || !reflect.DeepEqual(compiler.memory, freshMemory) {
		t.Error("Compiler is not clean")
	}
}
