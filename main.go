package main

import (
	"fmt"
	"os"
)

//go:generate go tool yacc -o parser.go -p parser parser.y

func main(){
	p, err := parse("def main(){ print 5; return 1; }")
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	ops, start, err := gen(p)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	s, err := dump(ops, start)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	fmt.Println(s)
	fmt.Println("Program Output")

	run(ops, start, os.Stdout)
}

func parse(prog string) (*SL, error) {
	l := &lex{ s : prog }
	r := parserParse(l)
	if r == -1 {
		return nil, fmt.Errorf("Unknown parser error encountered")
	}

	return l.result, nil
}