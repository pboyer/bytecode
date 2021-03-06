package main

import (
	"testing"
	"bytes"
)

func TestAST(t *testing.T){
	progs := []struct{
		prog *BlockS
		result string
	}{
		{
			makeAST(
				[]S {
					&PrintS{
						&BinOpE{ ADD, &IntE{ 3 }, &IntE{ 2 } },
					},
					&RetS{},
				},
				[]*FDefS{},
			),
			"5",
		},
		{
			makeAST(
				[]S {
					&PrintS{
						&BinOpE{ SUB, &IntE{ 2 }, &IntE{ 3 } },
					},
					&RetS{},
				},
				[]*FDefS{},
			),
			"-1",
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
					&RetS{},
				},
				[]*FDefS{
					&FDefS{
						name : "add2",
						args : []string{ "a" },
						body : &BlockS{
							list : []S {
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
					&RetS{},
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
					&RetS{},
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
					&RetS{},
				},
				[]*FDefS{
					&FDefS{
						name : "add2P",
						args : []string{"a" },
						body : &BlockS{
							list : []S{
								&RetS{
									&CallE{"add2", []E{ &IdE{ "a" } } },
								},
							},
						},
					},
					&FDefS{
						name : "add2",
						args : []string{"a" },
						body : &BlockS{
							list : []S{
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
		{
			makeAST(
				[]S{
					&RetS{
						&CallE{"printLocals", []E{}  },
					},
				},
				[]*FDefS{
					&FDefS{
						name : "printLocals",
						args : []string{},
						body : &BlockS{
							list : []S{
								&VDefS{
									name : "foo",
									rhs : &IntE{ 99 },
								},
								&VDefS{
									name : "bar",
									rhs : &IntE{ 55 },
								},
								&VDefS{
									name : "baz",
									rhs : &IntE{ 66 },
								},
								&PrintS{
									&IdE{"foo" },
								},
								&PrintS{
									&IdE{"bar" },
								},
								&PrintS{
									&IdE{"baz" },
								},
								&RetS{},
							},
						},
					},
				},
			),
			"995566",
		},
		{
			makeAST(
				[]S{
					&VDefS{
						name : "foo",
						rhs : &IntE{ 99 },
					},
					&RetS{
						&CallE{"printArgsLocals", []E{ &IdE{ "foo" }}  },
					},
				},
				[]*FDefS{
					&FDefS{
						name : "printArgsLocals",
						args : []string{ "a" },
						body : &BlockS{
							list : []S{
								&VDefS{
									name : "foo",
									rhs : &IntE{ 99 },
								},
								&PrintS{
									&BinOpE{ADD, &IdE{ "a" }, &IdE{ "foo" } },
								},
								&RetS{},
							},
						},
					},
				},
			),
			"198",
		},
		{
			makeAST(
				[]S{
					&IfS{
						test : &IntE{ 0 },
						tb : &BlockS{
							list : []S{
								&PrintS{&IntE{99 } },
							},
						},
						fb : &BlockS{
							list : []S{
								&PrintS{&IntE{55 } },
							},
						},
					},
					&RetS{},
				},
				[]*FDefS{},
			),
			"55",
		},
		{
			makeAST(
				[]S{
					&IfS{
						test : &IntE{ 1 },
						tb : &BlockS{
							list : []S{
								&PrintS{&IntE{99 } },
							},
						},
						fb : &BlockS{
							list : []S{
								&PrintS{&IntE{55 } },
							},
						},
					},
					&RetS{},
				},
				[]*FDefS{},
			),
			"99",
		},
		{
			makeAST(
				[]S{
					&PrintS{
						&CallE{"fibo", []E{ &IntE{ 8 } }  },
					},
					&RetS{},
				},
				[]*FDefS{
					&FDefS{
						name : "fibo",
						args : []string{ "a" },
						body : &BlockS{
							list : []S{
								&IfS{
									test: &BinOpE{ LEQ, &IdE{ "a" }, &IntE{ 1 } },
									tb : &BlockS{
										list : []S{
											&RetS{ &IntE{ 1 } },
										},
									},
								},
								&RetS{
									&BinOpE{ADD,
										&CallE{
											"fibo",
											[]E{
												&BinOpE{SUB, &IdE{ "a" }, &IntE{ 1 } },
											},
										},
										&CallE{
											"fibo",
											[]E{
												&BinOpE{SUB, &IdE{ "a" }, &IntE{ 2 } },
											},
										},
									},
								},
							},
						},
					},
				},
			),
			"34",
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