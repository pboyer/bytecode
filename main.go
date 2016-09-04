package main

import (
	"fmt"
	"os"
)

func main(){
	p := makeAST(
		[]S{
			&PrintS{
				&CallE{"add2P", []E{ &IntE{ 5 } }  },
			},
			&RetS{},
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

	ops, start, err := gen(p)
	if err != nil {
		fmt.Printf("Error : %v", err)
	}

	run(ops, start, os.Stdout)
}
