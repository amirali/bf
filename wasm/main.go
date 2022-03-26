package main

import (
	"fmt"
	"syscall/js"

	"github.com/amirali/bf"
)

func jsBrainfuckNewCompile(this js.Value, args []js.Value) (result interface{}) {
	program := args[0].String()
	input := []byte(args[1].String())

	compiler := bf.NewCompiler()

	err := compiler.Compile(program)
	if err != nil {
		return err.Error()
	}

	err = compiler.Execute(input)
	if err != nil {
		return err.Error()
	}

	return compiler.Output()
}

func main() {
	c := make(chan bool)
	fmt.Println("Brainfuck compiler loaded")
	js.Global().Set("execute_brainfuck", js.FuncOf(jsBrainfuckNewCompile))
	<-c
}
