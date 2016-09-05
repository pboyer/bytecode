package main

import (
	"fmt"
	"os"
)

//go:generate go tool yacc -o parser.go parser.y

func main(){
	p, err := parse(`
def foo(){
	var a = 2;
	return a >= 1;
}

def main(){
	print foo();
	return;
}`)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	ops, start, err := gen(p)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	s, err := dump(ops, start)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Println(s)
	fmt.Println("Program Output")

	run(ops, start, os.Stdout)
}

func parse(prog string) (*SL, error) {
	l := &lex{ s : prog }
	r := yyParse(l)
	fmt.Println(r)
	if r != 0 {
		return nil, fmt.Errorf("Unknown parser error encountered")
	}

	return l.result, nil
}