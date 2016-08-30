package main

import (
	"fmt"
)

func main(){
//	ops := []op {
//		op{ PUSH, 2 },
//		op{ PUSH, 1 },
//		op{ BIN_OP, ADD },
//		op{ PRINT, 0 },
//	}

	prog := &SL {
		ss : []S {
			&PrintS{
				&BinOpE{ ADD, &IntE{ 2 }, &IntE{ 2 } },
			},
		},
	}

	ops, start, err := gen(prog)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	run(ops, start)
}