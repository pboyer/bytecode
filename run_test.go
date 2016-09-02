package main

import (
	"testing"
	"bytes"
)

func TestAST(t *testing.T){
	progs := []struct{
		prog *SL
		result string
	}{
		{
			makeAST(
				[]S {
					&PrintS{
						&BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
					},
				},
				[]*FDefS{},
			),
			"5",
		},
		{
			makeAST(
				[]S{
					&PrintS{
						&CallE{"add2", []E{ &IntE{ 12 } }  },
					},
					&PrintS{
						&CallE{"add2", []E{ &IntE{ 5 } }  },
					},
				},
				[]*FDefS{
					&FDefS{
						name : "add2",
						args : []string{ "a" },
						body : &SL{
							ss : []S {
								&RetS{
									&BinOpE{ ADD, &IdE{ "a" }, &IntE{ 2 } },
								},
							},
						},
					},
				},
			),
			"147",
		},
		{
			makeAST(
				[]S{
					&VDefS{
						name : "foo",
						rhs : &IntE{ 2 },
					},
					&PrintS{
						&IdE{ "foo" },
					},
				},
				[]*FDefS{},
			),
			"2",
		},
		{
			makeAST(
				[]S {
					&VDefS {
						name : "foo",
						rhs : &BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
					},
					&PrintS {
						&IdE{ "foo" },
					},
				},
				[]*FDefS{},
			),
			"5",
		},
		{
			makeAST(
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
			),
			"7",
		},
	}

	for i, p := range progs {
		ops, start, err := gen(p.prog)
		if err != nil {
			t.Fatalf("Error at program %v : %v", i, err)
		}

		b := &bytes.Buffer{}

		run(ops, start, b)

		if b.String() != p.result {
			t.Fatalf("Unexpected result at program %v : expected %v, got %v", i, p.result, b.String())
		}
	}
}