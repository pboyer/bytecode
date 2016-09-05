.PHONY : build

build:
	go build 

run:
	go run `ls | grep -v _test | grep .go | tr '\n' ' '`

test:
	go test