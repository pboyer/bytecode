package main

import (
	"fmt"
)

func main(){

//	prog := prog(
//		[]S {
//			&PrintS{
//				&BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
//			},
//		},
//		[]*FDefS{},
//	)

//	prog2 := prog(
//		[]S{
//			&PrintS{
//				&CallE{"add2", []E{ &IntE{ 12 } }  },
//			},
//			&PrintS{
//				&CallE{"add2", []E{ &IntE{ 5 } }  },
//			},
//		},
//		[]*FDefS{
//			&FDefS{
//				name : "add2",
//				args : []string{ "a" },
//				body : &SL{
//					ss : []S {
//						&RetS{
//							&BinOpE{ ADD, &IdE{ "a" }, &IntE{ 2 } },
//						},
//					},
//				},
//			},
//		},
//	)

//	prog3 := prog(
//		[]S{
//			&VDefS{
//				name : "foo",
//				rhs : &IntE{ 2 },
//			},
//			&PrintS{
//				&IdE{ "foo" },
//			},
//		},
//		[]*FDefS{},
//	)

//	prog4 := prog(
//		[]S {
//			&VDefS {
//				name : "foo",
//				rhs : &BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
//			},
//			&PrintS {
//				&IdE{ "foo" },
//			},
//		},
//		[]*FDefS{},
//	)

	prog5 := prog(
		[]S{
			&PrintS{
				&CallE{"add2P", []E{ &IntE{ 5 } }  },
			},
		},
		[]*FDefS{
			&FDefS{
				name : "add2P",
				args : []string{"a" },
				body : &SL{
					ss : []S{
						&RetS{
							&CallE{"add2", []E{ &IdE{ "a" } } },
						},
					},
				},
			},
			&FDefS{
				name : "add2",
				args : []string{"a" },
				body : &SL{
					ss : []S{
						&RetS{
							&BinOpE{ADD, &IdE{ "a" }, &IntE{ 2 } },
						},
					},
				},
			},
		},
	)

	ops, start, err := gen(prog5)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	d, err := dump(ops, start)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	fmt.Println( d )

	run(ops, start)
}

func prog(stmts []S, fdefs []*FDefS) *SL {

	// main must return
	if len(stmts) == 0 {
		stmts = append(stmts, &RetS{})
	} else if _, ok := stmts[len(stmts)-1].(*RetS); !ok {
		stmts = append(stmts, &RetS{})
	}

	// inject main
	prog := &SL {
		ss : []S {
			&FDefS{
				name : "main",
				body: &SL{
					ss : stmts,
				},
			},
		},
	}

	// and additional fdefs
	for _, s := range fdefs {
		prog.ss = append(prog.ss, s)
	}

	return prog
}
