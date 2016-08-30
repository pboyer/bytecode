package main

import(
	"fmt"
)

var ops []op
var start int

func gen(program *SL) ([]op, int, error) {
	ops = make([]op, 0)
	err := genInt(program, newEnv(nil), true)
	if err != nil {
		return nil, -1, err
	}
	ops = append(ops, op{ HALT, -1 })
	return ops, start, nil
}

func genInt(n N, e *env, isGlobal bool) error {
	switch t := n.(type) {
	case *SL:
		// first pass - handle declarations
		var count int
		if isGlobal {
			count = 0
		} else {
			numParams := 1
			count = numParams
		}
		for _, s := range t.ss {
			switch ti := s.(type) {
			case *VDefS:
				// local
				e.data[ti.name] = &symbol{ count, "local", ti }
				count++
			case *FDefS:
				if !isGlobal {
					return fmt.Errorf("Nested function not allowed %v", ti.name)
				}
				sym := &symbol{len(ops), "func", ti}

				e.data[ti.name] = sym

				locals := make(map[string]*VDefS)
				for _, s := range ti.body.ss {
					if def, ok := s.(*VDefS); ok {
						if _, ok := locals[def.name]; !ok {
							locals[def.name] = def
						}
					}
				}
				ti.locals = locals

				// make new env with parameters
				et := newEnv(e)
				et.data[ti.param] = &symbol{ 0, "local", nil }

				err := genInt(ti.body, et, false)
				if err != nil {
					return err
				}
			}
		}

		if isGlobal {
			start = len(ops)
		}

		for i, s := range t.ss {
			switch ti := s.(type) {
			case *VDefS:
				if isGlobal {
					ops = append(ops, op{ PUSH, 0 })
				}
				t.ss[i] = &AssignS{ti.name, ti.rhs }
			}
		}

		for _, s := range t.ss {
			err := genInt(s, e, false)
			if err != nil {
				return err
			}
		}
	case *VDefS, *FDefS:
	case *IntE:
		ops = append(ops, op{ PUSH, t.val })
	case *PrintS:
		err := genInt(t.e, e, false)
		if err != nil {
			return err
		}
		ops = append(ops, op{ PRINT, -1 })
	case *BinOpE:
		err := genInt(t.lhs, e, false)
		if err != nil {
			return err
		}
		err = genInt(t.rhs, e, false)
		if err != nil {
			return err
		}
		ops = append(ops, op{ BIN_OP, int(t.op) })
	case *IdE:
		sym, ok := e.lookup(t.name)
		if !ok {
			return fmt.Errorf("Could not find %v in env", t.name)
		}

		switch sym.symbolType {
		case "local":
			ops = append(ops, op{ LOAD, sym.pos })
		default:
			return fmt.Errorf("Non-locals not allowed %v", t.name)
		}
	case *AssignS:
		// compute the expression, pushing it onto the stack
		err := genInt(t.rhs, e, false)
		if err != nil {
			return err
		}

		// what is the position of the id in the frame?
		sym,ok := e.lookup(t.lhs) // TODO globals
		if !ok {
			return fmt.Errorf("Could not find %s in environment", t.lhs)
		}

		// store the value of the expression in the appropriate position
		ops = append(ops, op{ STO, sym.pos })
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
		ops = append(ops, op{ PUSH_FP, -1 })

		// store space for return value
		ops = append(ops, op{ PUSH, 0 })

		// store the num args, locals for function
		ops = append(ops, op{ PUSH, len(fd.locals) + 1 })

		// push all of the args
		// TODO for now only one
		err := genInt(t.arg, e, false)
		if err != nil {
			return err
		}

		// make space for locals on stack
		for _ = range fd.locals {
			ops = append(ops, op{ PUSH, 0 })
		}

		// store the return address
		ops = append(ops, op{ PUSH_IP, -1 })

		// call function
		ops = append(ops, op{ JMP, sym.pos })
	case *RetS:
		if isGlobal {
			return fmt.Errorf("Unexpected return statement")
		}

		// compute the return value
		if t.rhs != nil {
			err := genInt(t.rhs, e, false)
			if err != nil {
				return err
			}

			// store in return value pos
			ops = append(ops, op{ STO, 1 })
		}

		// restore the fp from the stack
		ops = append(ops, op{ RET, -1 })

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
//ops = append(ops, bop)
//
//// emit true branch
//cgen(n.tb, e, data)
//
//// after completing tb, jmp to after fb
//tbd := op{ code : JMP }
//ops = append(ops, tbd)
//
//bop.op1 = len(ops) // set the branch address
//
//// emit the false branch
//cgen(n.fb, e, data)
//
//tbd.op1 = len(ops) // set the true branch jmp
//

