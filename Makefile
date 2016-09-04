.PHONY : build

build:
	go build 

run:
	go run env.go gen.go main.go opcode_string.go binop_string.go ast.go run.go

test:
	go test