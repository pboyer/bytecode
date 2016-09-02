package main

import(
	"fmt"
	"bytes"
)

//go:generate stringer -type=opCode
//go:generate stringer -type=binOp

type binOp int

const (
	ADD binOp = iota
	SUB
	DIV
	MUL
)

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

func pop() (int) {
	if len(stack) == 0 {
		panic("pop called on empty stack")
	}
	v := stack[len(stack)-1]
	stack = stack[0:len(stack)-1]
	return v
}

func push(val int) {
	stack = append(stack, val)
}

func get(i int) int {
	if i >= len(stack) {
		panic(fmt.Errorf("Index out of range:%v\n", stack))
	}
	return stack[i]
}

func dump(ops []op, pos int) (string, error) {
	buf := &bytes.Buffer{}

	var err error

	for i, op := range ops {
		if pos >= 0 && i == pos {
			_, err = fmt.Fprintf(buf, "\t%d:\t%v <- \n", i, op )
		} else {
			_, err = fmt.Fprintf(buf, "\t%d:\t%v\n", i, op )
		}

		if err != nil {
			return "", err
		}
	}
	return string(buf.Bytes()), nil
}

func run(ops []op, pc int) {

	defer func(){
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
		if pc >= len(ops){
			break
		}

		op := ops[pc]

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
		case PRINT:
			var r int
			r = pop()
			fmt.Println("PRINT", r)
			pc++
		case PUSH_IP:
			push(pc+1)
			pc++
		case BIN_OP:
			var r1, r2 int
			r1 = pop()
			r2 = pop()

			switch binOp(op.op1) {
			case ADD:
				push(r1+r2)
			case SUB:
				push(r1+r2)
			case DIV:
				push(r1/r2)
			case MUL:
				push(r1*r2)
			}
			pc++
		case PUSH_FP:
			fpt := fp
			fp = len(stack)
			push(fpt)
			pc++
		case RET:
			// 0 args
			// 1 locals
			// 2 varCount
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
			var r1,pos int
			if op.op1 >= 0 {
				pos = op.op1
				varCount := stack[fp-1]
				pos = fp-1-varCount+pos
			} else {
				pos = pop()
			}
			r1 = pop()
			stack[pos] = r1
			pc++
		case LOAD:
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

