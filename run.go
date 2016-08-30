package main

import(
	"fmt"
)

type BinOp int

const (
	ADD BinOp = iota
	SUB
	DIV
	MUL
)

type Op int

const (
	PRINT Op = iota
	PUSH
	POP
	RET
	STO
	LOAD
	BIN_OP
	PUSH_IP // store ip+1 on stack
	PUSH_FP
	JMP
)

type op struct {
	code Op
	op1 int
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

func run(ops []op, pc int){
	stack = []int{ 0, 0, 0 }

	fp := 0

	for {
		if pc >= len(ops){
			break
		}

		op := ops[pc]

		switch op.code {
		case PRINT:
			var r int
			r = pop()
			fmt.Println(r)
			pc++
		case PUSH_FP:
			push(fp)
			pc++
		case PUSH_IP:
			push(pc+1)
			pc++
		case PUSH:
			push(op.op1)
			pc++
		case POP:
			_ = pop()
			pc++
		case BIN_OP:
			var r1, r2 int
			r1 = pop()
			r2 = pop()

			switch op.op1 {
			case ADD:
				push(r1+r2)
			case SUB:
				push(r1+r2)
			case DIV:
				push(r1/r2)
			case MULT:
				push(r1*r2)
			}
			pc++
		case PUSH_FP:
			fpt := fp
			fp = len(stack)
			push(fpt)
		case RET:
			sp := stack[fp]
			rv := stack[fp + 1]
			numArgsLocals := stack[fp+2]
			pc = stack[fp + numArgsLocals + 3]
			for len(stack) > sp {
				pop()
			}
			push(rv)
		case STO:
			var r1,pos int
			if op.op1 >= 0 {
				pos = op.op1+3
			} else {
				pos = pop()
			}

			r1 = pop()
			stack[pos] = r1
			pc++
		case LOAD:
			pos := op.op1
			push(stack[fp+pos+2]) // +2 accounts for the fp, returnAddr, returnVal
			pc++
		default:
			panic("unknown opcode")
		}
	}
}

