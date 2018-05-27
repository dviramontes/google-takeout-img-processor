PHONY: build
build:
	go build -o proc main.go

PHONY: run
run:
	root=${root} dest=${dest} ./proc
