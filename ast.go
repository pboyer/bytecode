package main

func makeAST(stmts []S, fdefs []*FDefS) *SL {

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

type SL struct {
	ss []S
}

func (s *SL) impleN(){}

type PrintS struct {
	e E
}

type AssignS struct {
	name string
	rhs  E
}

type IfS struct {
	test E
	tb *SL
	fb *SL
}

type VDefS struct {
	name string
	rhs E
}

type FDefS struct {
	name   string
	args   []string
	body   *SL
	locals map[string]*VDefS
}

type RetS struct {
	rhs E
}

func (s *PrintS) impleN(){}
func (s *AssignS) impleN(){}
func (s *IfS) impleN(){}
func (s *FDefS) impleN(){}
func (s *VDefS) impleN(){}
func (s *RetS) impleN(){}

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

