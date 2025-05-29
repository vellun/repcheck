CMD_DIR=./cmd

BINARY=repcheck

build:
	go build -o $(BINARY) $(CMD_DIR)
