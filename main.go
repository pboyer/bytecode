package main

import (
	"fmt"
	"os"
)

//go:generate go tool yacc -o parser.go parser.y

func main() {
	p, err := parse(`
def fibo(n){
	if (n <= 1){
		return 1;
	}
	return fibo(n-1) + fibo(n-2);
}

def main(){
	print fibo(6);
	print fibo(7);
	print fibo(8);
	return;
}`)
	if err != nil {
		return
	}

	err = check(p)

	ops, start, err := gen(p)
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		return
	}

	// s, err := dump(ops, start)
	// if err != nil {
	// 	fmt.Printf("Error : %v\n", err)
	// 	return
	// }

	// fmt.Println(s)

	run(ops, start, os.Stdout)
}
