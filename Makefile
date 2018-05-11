GOPATH=${PWD}
MAIN=main.go

all: setup build

setup:
	export GOPATH=${PWD}

build: setup
	go build

run: setup
	go run ${MAIN}

clean: setup
	go clean
