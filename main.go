package main

import (
	"fmt"
	"os"
)

func main(){
	p := makeAST(
		[]S{
			&IfS{
				test : &IntE{ 0 },
				tb : &SL{
					ss : []S{
						&PrintS{&IntE{99 } },
					},
				},
				fb : &SL{
					ss : []S{
						&PrintS{&IntE{55 } },
					},
				},
			},
			&RetS{},
		},
		[]*FDefS{},
	)

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
