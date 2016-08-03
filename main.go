package main

import (
    "fmt"
)

const (
    ADD = iota
    SUB
)

const (
    PRINT = iota
    PUSH
    PUSHFP // push fp onto stack
    FP // set fp from the stack
	PUSHPC // store ip+1 on stack
    STO
	LD
    POP
    BOP
    CALL
    JMP
)

type op struct {
    code int
    op1 int
}

func pop(stack []int) (int,[]int) {
    if len(stack) == 0 {
        panic("Empty stack")
    }

    return stack[len(stack)-1], stack[:len(stack)-1]
}

func push(stack []int, val int) []int {
    return append(stack, val)
}

func main(){
    sp := 0
    fp := 0
    pc := 0
    
    stack := make([]int, 0)
    ops := []op {
        op{ PUSH, 2 },
        op{ PUSH, 1 },
        op{ BOP, ADD },
        op{ PRINT, 0 },
    }

    for {
        if pc >= len(ops){
            break
        }
       
        op := ops[pc]

        switch op.code {
        case PRINT:
            var r int
            r,stack = pop(stack)
            fmt.Println(r)
            pc++
        case PUSHFP:
	        stack = push(stack, fp)
	        pc++
        case PUSHPC:
	        stack = push(stack, pc+1)
	        pc++
        case PUSH:
            stack = push(stack, op.op1)
            pc++
        case POP:
            _,stack = pop(stack)
            pc++
        case BOP:
            var r1, r2 int
            r1,stack = pop(stack)
            r2,stack = pop(stack)
            
            switch op.op1 {
            case ADD:
                stack = push(stack, r1+r2)
            case SUB:
                stack = push(stack, r1+r2)
            }
            pc++
        case STO:
	        var r1,pos int
	        pos,stack = pop(stack)
	        r1,stack = pop(stack)

	        stack[pos] = r1
	        pc++
        case LD:
	        pos := op.op1
	        stack = push(stack, stack[fp+pos])
	        pc++
        case CALL:
	        pc = op.op1
	        fp = len(stack)
		default:
	        panic("unknown opcode")
        }
    }
}

func cgen(n N, e *env, ops []op)  {
    switch t := n.(type) {
    case *SL:
	    // first pass - handle declarations
	    count := 0
	    for i, s := range SL.ss {
		    switch t := s.(type) {
		    case *VDefS:
			    // local
			    e[t.name] = symbol{count, "local" }
			    count++

			    // we'll handle as an assignment
			    SL.ss[i] = &AssignS{t.name, t.rhs }
		    case *FDefS:
			    fd := symbol{len(ops), "func" }

			    e.data[t.name] = fd

			    en := newEnv(nil) // TODO include globals

			    en.data[t.name] = fd // for recursion
			    en.data[t.param] = symbol{2, "local" }

			    cgen(t.body, en, ops)
		    }
	    }

	    for _, s := range SL.ss {
		    cgen(s, e, ops)
	    }
    case *VDefS, *FDefS:
    case *IntE:
	    ops = append(ops, op{PUSH, t.val })
    case *AddE:
	    cgen(t.lhs, e, ops)
	    cgen(t.rhs, e, ops)
	    ops = append(ops, op{BOP, ADD })
    case *SubE:
	    cgen(t.rhs, e, ops)
	    cgen(t.lhs, e, ops)
	    ops = append(ops, op{BOP, SUB })
    case *IdE:
	    sym, ok := e.data[t.name]
	    if !ok {
		    panic("Could not find " + t.name + " in env")
	    }

	    switch sym {
	    case "local":
		    // compute the position, get the value from the position, place on stack
		    panic("TODO")
	    }
    case *AssignS:
        // compute the expression, pushing it onto the stack
        cgen(t.rhs, e, ops)
               
        // what is the position of the id in the frame?
        pos := e[t.lhs]

        // push the frame pointer onto the stack
        ops = append(ops, op{ code : PUSHFP })

        // push the offset
        ops = append(ops, op{ PUSH, pos })
        
        // add these, yielding the position on the stack
        ops = append(ops, op{ BOP, ADD })

        // store the value of the expression
        ops = append(ops, op{ STO })
    case *RetS:
        // restore the frame pointer from the stack
	    ops = append(ops, op{ FP })

	    // compute the return value
	    cgen(t.rhs, e, ops)
    case *CallE:
        // push all of the args in reverse order TODO (there's only one!)
        cgen(t.arg, e, ops)

	    // make space for locals on stack
		for _,v := range e.data {
			if v.symbolType == "local" {
				ops = append(ops, op{ PUSH, 0 })
			}
		}

	    // store the return address in the return address register
	    ops = append(ops, op{ code : PUSHPC })

	    // store the current frame pointer on the stack
	    ops = append(ops, op{ code : PUSHFP })

        // goto function def!
        ops = append(ops, op{ CALL, e.lookup(t.name) })
    }
}

// --------------- fp
// args
// locals
// return address
// store old fp


//case *IfS:
//// compute test
//cgen(n.test, e, ops, data)
//
//// create op reference
//bop := op{ code : BIF }
//
//// branch if false
//ops = append(ops, bop)
//
//// emit true branch
//cgen(n.tb, e, ops, data)
//
//// after completing tb, jmp to after fb
//tbd := op{ code : JMP }
//ops = append(ops, tbd)
//
//bop.op1 = len(ops) // set the branch address
//
//// emit the false branch
//cgen(n.fb, e, ops, data)
//
//tbd.op1 = len(ops) // set the true branch jmp
//

type N interface {
	impleN()
}

type E interface {
	impleE()
}

type S interface {
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
	lhs string
	rhs E
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
	name string
	param string
	body *SL
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

type AddE struct {
	rhs, lhs E
}

type SubE struct {
	rhs, lhs E
}

type CallE struct {
	name string
	arg E
}

type IdE struct {
	name string
}

type IntE struct {
	val int
}

func (s *SubE) impleN(){}
func (s *AddE) impleN(){}
func (s *CallE) impleN(){}
func (s *IdE) impleN(){}
func (s *IntE) impleN(){}

func (s *SubE) impleE(){}
func (s *AddE) impleE(){}
func (s *CallE) impleE(){}
func (s *IdE) impleE(){}
func (s *IntE) impleE(){}
/*
const (
    PRINT = iota
    PUSH
    RA // get the return address reg and put it in on the stack
    FP // place the fp register onto the stack, do not modify the fp
    MOV // MOV dest sourcevalue i.e. dest is on top of stack
    CALL // call the function - assumes the frame pointer, return address,
    RET // return from the function including restoring the frame pointer, pushing the return value and jmp to ra
    LD
    POP
    BOP
    CALL
    JMP
    BIF // branch if != 0
)

type op struct {
    code int
    op1 int
}
*/


// TODO no globals for now

type symbol struct {
	pos int
	symbolType string
}

type env struct {
	parent *env
	data map[string]symbol
}

func newEnv(parent *env) *env {
	return &env{
		parent : parent,
		data : make(map[string]symbol),
	}
}

func (e *env) lookup(id string) (symbol,bool) {
	s,ok := e.data[id]
	if ok {
		return s,true
	}

	if e.parent == nil {
		return nil, false
	}

	return e.parent.lookup(id)
}

