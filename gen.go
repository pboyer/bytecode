package main

import(
	"fmt"
)

var ops []op
var unresolvedJmps map[string][]int
var start int

func gen(program *SL) ([]op, int, error) {
	ops = make([]op, 0)
	unresolvedJmps = make(map[string][]int)

	globalEnv := newEnv(nil)
	err := genInt(program, globalEnv, true, 0)
	if err != nil {
		return nil, -1, err
	}

	err = resolve(globalEnv)
	if err != nil {
		return nil, -1, err
	}

	// goto main
	start = len(ops)
	genInt(&CallE{
		name : "main",
		args : []E{},
	}, globalEnv, false, 0)

	ops = append(ops, op{ HALT, -1 })
	return ops, start, nil
}

func resolve(globals *env) error {
	for k, v := range unresolvedJmps {
		for _, i := range v {
			opToFix := ops[i]
			if opToFix.code != JMP {
				return fmt.Errorf("Non jmp op found unresolved at position %v", i)
			}

			if opToFix.op1 != -1 {
				return fmt.Errorf("Already resolved jmp op found at position %v", i)
			}

			if sym, ok := globals.data[k]; ok {
				ops[i].op1 = sym.pos
			} else {
				return fmt.Errorf("Unresolved instruction at %v for symbol", i, k)
			}
		}
	}
	return nil
}

func genInt(n N, e *env, isGlobal bool, argCount int) error {
	switch t := n.(type) {
	case *SL:
		// discover function definition
		for _, s := range t.ss {
			switch ti := s.(type) {
			case *FDefS:
				if !isGlobal {
					return fmt.Errorf("Nested function not allowed %v", ti.name)
				}
				// we don't know code positions yet, we'll need to resolve later
				sym := &symbol{-1, "func", ti}
				e.data[ti.name] = sym
			}
		}

		// compile function definitions
		count := argCount
		for _, s := range t.ss {
			switch ti := s.(type) {
			case *VDefS:
				e.data[ti.name] = &symbol{ count, "local", ti }
				count++
			case *FDefS:
				e.data[ti.name].pos = len(ops)

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
				for i, arg := range ti.args {
					et.data[arg] = &symbol{ i, "arg", nil }
				}

				err := genInt(ti.body, et, false, len(ti.args))
				if err != nil {
					return err
				}
			}
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
			err := genInt(s, e, false, argCount)
			if err != nil {
				return err
			}
		}
	case *VDefS, *FDefS:
	case *IntE:
		ops = append(ops, op{ PUSH, t.val })
	case *PrintS:
		err := genInt(t.e, e, false, argCount)
		if err != nil {
			return err
		}
		ops = append(ops, op{ PRINT, -1 })
	case *BinOpE:
		err := genInt(t.lhs, e, false, argCount)
		if err != nil {
			return err
		}
		err = genInt(t.rhs, e, false, argCount)
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
		case "local","arg":
			ops = append(ops, op{ LOAD, sym.pos })
		default:
			return fmt.Errorf("Non-locals not allowed %v", t.name)
		}
	case *AssignS:
		// compute the expression, pushing it onto the stack
		err := genInt(t.rhs, e, false, argCount)
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
		// 0 args
		// 1 locals
		// 2 varCount
		// 3 oldfp               FP
		// 4 return address
		// 5 return value        SP

		// look up the function
		sym, ok := e.lookupRec(t.name)
		if !ok {
			return fmt.Errorf("Could not find func %v in environment", t.name)
		}

		fd,ok := sym.node.(*FDefS)
		if !ok {
			return fmt.Errorf("%v is not a function", t.name)
		}

		// push all of the args
		for _, arg := range t.args {
			err := genInt(arg, e, false, argCount)
			if err != nil {
				return err
			}
		}

		// make space for locals on stack
		for _ = range fd.locals {
			ops = append(ops, op{ PUSH, 0 })
		}

		// store the num args, locals for function
		ops = append(ops, op{ PUSH, len(fd.locals) + len(fd.args) })

		// store the current frame pointer on the stack, implicitly sets frame pointer register to stack pointer
		ops = append(ops, op{ PUSH_FP, -1 })

		// store the return address
		ops = append(ops, op{ PUSH_IP, -1 })

		// call function
		ops = append(ops, op{ JMP, sym.pos })

		// if, we have yet to compile the function
		// store this code position for later
		if sym.pos == -1 {
			if _, ok := unresolvedJmps[t.name]; ok {
				unresolvedJmps[t.name] = append(unresolvedJmps[t.name], len(ops)-1)
			} else {
				unresolvedJmps[t.name] = []int{ len(ops)-1 }
			}
		}
	case *RetS:
		if isGlobal {
			return fmt.Errorf("Unexpected return statement")
		}

		// compute the return value
		if t.rhs != nil {
			err := genInt(t.rhs, e, false, argCount)
			if err != nil {
				return err
			}
		} else {
			// if no ret value, simply push 0
			ops = append(ops, op{ PUSH, 0 })
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

