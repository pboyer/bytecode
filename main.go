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
//				&CallE{ "add2", &IntE{ 5 }  },
//			},
//		},
//	}

//	prog3 := &SL {
//		ss : []S {
//			&VDefS {
//				name : "foo",
//				rhs : &IntE{ 2 },
//			},
//			&PrintS {
//				&IdE{ "foo" },
//			},
//		},
//	}

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

	prog5 := &SL {
		ss : []S {
			&FDefS{
				name : "add2P",
				param : "a",
				body : &SL{
					ss : []S {
						&RetS{
							&CallE{ "add2", &IdE{ "a" } },
						},
					},
				},
			},
			&FDefS{
				name : "add2",
				param : "a",
				body : &SL{
					ss : []S {
						&RetS{
							&BinOpE{ ADD, &IdE{ "a" }, &IntE{ 2 } },
						},
					},
				},
			},
			&PrintS {
				&CallE{ "add2P", &IntE{ 5 }  },
			},
		},
	}

	ops, start, err := gen(prog5)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	d, err := dump(ops, -1)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	fmt.Println( d )

	run(ops, start)
}