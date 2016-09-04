package main

import (
	"fmt"
	"os"
)

//go:generate go tool yacc -o parser.go -p parser parser.y

func main(){

	parserParse(&lex{})

	ops, start, err := gen(p)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	s, err := dump(ops, start)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	fmt.Println(s)

	run(ops, start, os.Stdout)
}
