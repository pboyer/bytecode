package main

import(
	"fmt"
)

var ops []op
var start int

func gen(program *SL) ([]op, int, error) {
	err := genInt(program, newEnv(nil), true)
	if err != nil {
		return nil, -1, err
	}
	return ops, start, nil
}

func genInt(n N, e *env, isGlobal bool) error {
	switch t := n.(type) {
	case *SL:
		// first pass - handle declarations
		count := 0
		for i, s := range SL.ss {
			switch t := s.(type) {
			case *VDefS:
				// local
				e[t.name] = symbol{ count, "local" }
				count++

				// we'll handle as an assignment
				SL.ss[i] = &AssignS{t.name, t.rhs }
			case *FDefS:
				if !isGlobal {
					return fmt.Errorf("Nexted function %v", t.name)
				}
				sym := symbol{len(ops), "func", t}

				e.data[t.name] = sym

				locals := make(map[string]*VDefS)
				for _, s := range t.body.ss {
					if def, ok := s.(*VDefS); ok {
						if _, ok := locals[def.name]; !ok {
							locals[def.name] = def
						}
					}
				}
				t.locals = locals

				genInt(t.body, e, false)
			}
		}

		for _, s := range SL.ss {
			genInt(s, e, false)
		}
	case *VDefS, *FDefS:
	case *IntE:
		ops = append(op{ PUSH, t.val })
	case *PrintS:
		genInt(t.e, e, false)
		ops = append(op{ PRINT })
	case *BinOpE:
		genInt(t.lhs, e, false)
		genInt(t.rhs, e, false)
		switch t.op {
		case ADD:
			ops = append(op{ BIN_OP, ADD })
		case SUB:
			ops = append(op{ BIN_OP, SUB })
		case MUL:
			ops = append(op{ BIN_OP, MUL })
		case DIV:
			ops = append(op{ BIN_OP, DIV })
		}

	case *IdE:
		sym, ok := e.lookup(t.name)
		if !ok {
			return fmt.Errorf("Could not find %v in env", t.name)
		}

		switch sym {
		case "local":
			ops = append(op{ LOAD, sym.pos })
		default:
			return fmt.Errorf("Non-locals not allowed %v", t.name)
		}
	case *AssignS:
		// compute the expression, pushing it onto the stack
		genInt(t.rhs, e, false)

		// what is the position of the id in the frame?
		sym,ok := e.lookup(t.lhs) // TODO globals
		if !ok {
			return fmt.Errorf("Could not find %s in environment", t.lhs)
		}

		// store the value of the expression in the appropriate position
		ops = append(op{ STO, sym.pos })
	case *CallE:
		// ...
		// --------------- fp
		// oldfp
		// return value
		// numArgs+numLocals
		// args
		// locals
		// return address
		// --------------- sp

		// look up the function
		sym, ok := e.lookupRec(t.name)
		if !ok {
			return fmt.Errorf("Could not find func %v in environment", t.name)
		}

		fd,ok := sym.node.(*FDefS)
		if !ok {
			return fmt.Errorf("%v is not a function", t.name)
		}

		// store the current frame pointer on the stack, implicitly sets frame pointer register to stack pointer
		ops = append(op{ code : PUSH_FP })

		// store space for return value
		ops = append(op{ PUSH, 0 })

		// store the num args, locals for function
		ops = append(op{ code : PUSH, op1 : len(fd.locals) + 1 })

		// make new env
		e = newEnv(e)

		// push all of the args in reverse order TODO for now only one
		genInt(t.arg, e, false)
		e.data[fd.param]

		// make space for locals on stack
		for _ := range fd.locals {
			ops = append(op{ PUSH, 0 })
		}

		// store the return address in the return address
		ops = append(op{ code : PUSH_IP })

		// call function
		ops = append(op{ JMP, sym.pos })
	case *RetS:
		if isGlobal {
			return fmt.Errorf("Unexpected return statement")
		}

		// compute the return value
		if t.rhs != nil {
			genInt(t.rhs, e, false)

			// store in return value pos
			ops = append(op{ STO, 1 })
		}

		// restore the ip from the stack and jump to return address
		ops = append(op{ LOAD, 2 })
		ops = append(op{ JMP })

		// restore the fp from the stack
		ops = append(op{ RET })

		e = e.parent
	}
	return nil
}

//case *IfS:
//// compute test
//cgen(n.test, e, data)
//
//// create op reference
//bop := op{ code : BIF }
//
//// branch if false
//ops = append(bop)
//
//// emit true branch
//cgen(n.tb, e, data)
//
//// after completing tb, jmp to after fb
//tbd := op{ code : JMP }
//ops = append(tbd)
//
//bop.op1 = len(ops) // set the branch address
//
//// emit the false branch
//cgen(n.fb, e, data)
//
//tbd.op1 = len(ops) // set the true branch jmp
//

