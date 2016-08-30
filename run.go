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
	return fmt.Sprintf("%v\t%v", o.code, o.op1)
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
			fmt.Println("Crash dump")
			fmt.Println(dump(ops, pc))
			fmt.Println(stack)
			panic(e)
		}
	}()

	stack = []int{ 0, 0, 0 }

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
			fmt.Println(r)
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
			sp := stack[fp]
			rv := stack[fp + 1]
			numArgsLocals := stack[fp+2]
			pc = stack[fp + numArgsLocals + 3]
			for len(stack) > sp {
				pop()
			}
			push(rv)
			pc++
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
			push(stack[fp+pos+3]) // +3 accounts for the fp, returnAddr, returnVal
			pc++
		case HALT:
			return
		default:
			panic(fmt.Errorf("unknown opcode %v", op))
		}
	}
}

