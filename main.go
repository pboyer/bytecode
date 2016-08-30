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

//	prog := &SL {
//		ss : []S {
//			&PrintS{
//				&BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
//			},
//		},
//	}
//
//	prog2 := &SL {
//		ss : []S {
//			&FDefS{
//				name : "add2",
//				param : "a",
//				body : &SL{
//					ss : []S {
//						&RetS{
//							&BinOpE{ ADD, &IdE{ "a" }, &IntE{ 2 } },
//						},
//					},
//				},
//			},
//			&PrintS {
//				&CallE{ "add2", &IntE{ 1 }  }
//			},
//		},
//	}

	prog3 := &SL {
		ss : []S {
			&VDefS {
				name : "foo",
				rhs : &IntE{ 2 },
			},
			&PrintS {
				&IdE{ "foo" },
			},
		},
	}

//	prog4 := &SL {
//		ss : []S {
//			&VDefS {
//				name : "foo",
//				rhs : &BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
//			},
//			&PrintS {
//				&IdE{ "foo" },
//			},
//		},
//	}

	ops, start, err := gen(prog3)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	d, err := dump(ops)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	fmt.Println( d )

	run(ops, start)
}