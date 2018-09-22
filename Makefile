GOPATH=${PWD}
MAIN=main.go

all: .setup build

.setup:
	eval "export GOPATH=${PWD}"

build: .setup
	go build -v

run: .setup
	go run ${MAIN}

clean: .setup
	go clean
