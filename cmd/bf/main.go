package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/amirali/bf"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	compiler := bf.NewCompiler()
	for {
		fmt.Print("> ")
		program, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error while reading program: %v", err)
		}

		err = compiler.Compile(program)
		if err != nil {
			log.Fatalln(err)
		}

		input := []byte{}
		inputCount := strings.Count(program, ",")
		if inputCount > 0 {
			fmt.Printf("(%d char input)>> ", inputCount)
			for ; inputCount > 0; inputCount-- {
				b, err := reader.ReadByte()
				if err != nil {
					log.Fatalf("Error while reading input: %v", err)
				}
				input = append(input, b)
			}
		}

		err = compiler.Execute(input)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(compiler.Output())

		err = compiler.Clear()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
