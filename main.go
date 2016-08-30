package main

func main(){
	ops := []op {
		op{ PUSH, 2 },
		op{ PUSH, 1 },
		op{ BIN_OP, ADD },
		op{ PRINT, 0 },
	}

	run(ops, 0)
}