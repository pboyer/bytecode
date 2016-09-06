package main

//go:generate stringer -type=binOp

func makeAST(stmts []S, fdefs []*FDefS) *BlockS {
	// inject main
	prog := &BlockS{
		list : []S {
			&FDefS{
				name : "main",
				body: &BlockS{
					list : stmts,
				},
			},
		},
	}

	// and additional fdefs
	for _, s := range fdefs {
		prog.list = append(prog.list, s)
	}

	return prog
}

type N interface {
	impleN()
}

type E interface {
	N
	impleE()
}

type S interface {
	N
	impleS()
}

type BlockS struct {
	list []S
}

type PrintS struct {
	e E
}

type AssignS struct {
	name string
	rhs  E
}

type IfS struct {
	test E
	tb S
	fb S
}

type VDefS struct {
	name string
	rhs E
}

type FDefS struct {
	name   string
	args   []string
	body   *BlockS
	locals map[string]*VDefS
}

type RetS struct {
	rhs E
}

func (s *BlockS) impleN(){}
func (s *PrintS) impleN(){}
func (s *AssignS) impleN(){}
func (s *IfS) impleN(){}
func (s *FDefS) impleN(){}
func (s *VDefS) impleN(){}
func (s *RetS) impleN(){}

func (s *BlockS) impleS(){}
func (s *PrintS) impleS(){}
func (s *AssignS) impleS(){}
func (s *IfS) impleS(){}
func (s *VDefS) impleS(){}
func (s *FDefS) impleS(){}
func (s *RetS) impleS(){}

type BinOpE struct {
	op       binOp
	rhs, lhs E
}

type binOp int

const (
	ADD binOp = iota
	SUB
	MUL
	DIV
	MOD
	GT
	LT
	EQ
	NEQ
	GEQ
	LEQ
	AND
	OR
)

type CallE struct {
	name string
	args []E
}

type IdE struct {
	name string
}

type IntE struct {
	val int
}

func (s *BinOpE) impleN(){}
func (s *CallE) impleN(){}
func (s *IdE) impleN(){}
func (s *IntE) impleN(){}

func (s *BinOpE) impleE(){}
func (s *CallE) impleE(){}
func (s *IdE) impleE(){}
func (s *IntE) impleE(){}
