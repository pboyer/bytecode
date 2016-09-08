package main

import (
	"bytes"
	"fmt"
	"io"
)

//go:generate stringer -type=opCode

type opCode int

const (
	PRINT opCode = iota
	PUSH
	POP
	RET
	STO
	LOAD
	BIN_OP
	PUSH_IP // store ip+1 on stack
	PUSH_FP
	JMP
	CJMP
	HALT
)

type op struct {
	code opCode
	op1  int
}

func (o op) String() string {
	if o.op1 >= 0 {
		return fmt.Sprintf("%v\t%v", o.code, o.op1)
	}
	return fmt.Sprintf("%v", o.code)
}

var stack []int

func pop() int {
	if len(stack) == 0 {
		panic("pop called on empty stack")
	}
	v := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	return v
}

func push(val int) {
	stack = append(stack, val)
}

func dump(ops []op, pos int) (string, error) {
	buf := &bytes.Buffer{}

	var err error

	for i, op := range ops {
		if pos >= 0 && i == pos {
			_, err = fmt.Fprintf(buf, "\t%d:\t%v <- \n", i, op)
		} else {
			_, err = fmt.Fprintf(buf, "\t%d:\t%v\n", i, op)
		}

		if err != nil {
			return "", err
		}
	}

	return string(buf.Bytes()), nil
}

func run(ops []op, pc int, writer io.Writer) {

	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Crash dump:")
			fmt.Println(dump(ops, pc))
			fmt.Println(stack)
			panic(e)
		}
	}()

	stack = []int{}

	fp := 0

	for {
		if pc >= len(ops) {
			break
		}

		op := ops[pc]

		// debug
		//		fmt.Println(op, "fp", fp, "pc", pc, stack)

		switch op.code {
		case PUSH:
			push(op.op1)
			pc++
		case POP:
			_ = pop()
			pc++
		case JMP:
			var p int
			if op.op1 >= 0 {
				p = op.op1
			} else {
				p = pop()
			}

			pc = p
		case CJMP:
			if pop() != 0 {
				pc++
			} else {
				pc = op.op1
			}
		case PRINT:
			var r int
			r = pop()
			fmt.Fprintf(writer, "%v\n", r)
			pc++
		case PUSH_IP:
			push(pc + 1)
			pc++
		case BIN_OP:
			var r1, r2 int
			r1 = pop()
			r2 = pop()

			switch binOp(op.op1) {
			case ADD:
				push(r1 + r2)
			case SUB:
				push(r1 - r2)
			case MUL:
				push(r1 * r2)
			case DIV:
				push(r1 / r2)
			case MOD:
				push(r1 % r2)
			case GT:
				if r1 > r2 {
					push(1)
				} else {
					push(0)
				}
			case LT:
				if r1 < r2 {
					push(1)
				} else {
					push(0)
				}
			case EQ:
				if r1 == r2 {
					push(1)
				} else {
					push(0)
				}
			case NEQ:
				if r1 != r2 {
					push(1)
				} else {
					push(0)
				}
			case GEQ:
				if r1 >= r2 {
					push(1)
				} else {
					push(0)
				}
			case LEQ:
				if r1 <= r2 {
					push(1)
				} else {
					push(0)
				}
			case AND:
				if (r1 != 0) && (r2 != 0) {
					push(1)
				} else {
					push(0)
				}
			case OR:
				if (r1 != 0) || (r2 != 0) {
					push(1)
				} else {
					push(0)
				}
			}
			pc++
		case PUSH_FP:
			fpt := fp
			fp = len(stack)
			push(fpt)
			pc++
		case RET:
			// Structure of the stack frame

			// 0 args
			// 1 locals
			// 2 varCount = len(args) + len(locals)
			// 3 oldfp               FP
			// 4 return address
			// 5 return value        SP

			// pop the return value
			rv := pop()

			// we'll need this offset
			varCount := stack[fp-1]

			newFp := stack[fp]

			retAd := stack[fp+1]

			// pop all remnants of the func call off
			sp := fp - 1 - varCount
			for len(stack) > sp {
				pop()
			}

			// restore old frame pointer
			fp = newFp

			// move to the return address
			pc = retAd

			// push the return value onto the stack
			push(rv)

			pc++
		case STO:
			var r1, pos int

			// the position on the stack is provided as operand
			if op.op1 >= 0 {
				pos = op.op1
				varCount := stack[fp-1]
				pos = fp - 1 - varCount + pos
				// the position is on the top of the stack
			} else {
				pos = pop()
			}
			r1 = pop()
			stack[pos] = r1
			pc++
		case LOAD:
			// the position on the stack is provided as operand
			pos := op.op1
			varCount := stack[fp-1]
			push(stack[fp-1-varCount+pos])
			pc++
		case HALT:
			return
		default:
			panic(fmt.Errorf("unknown opcode %v", op))
		}
	}
}
