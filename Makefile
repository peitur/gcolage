GOPATH=${PWD}
MAIN=main.go

all: .setup build

.setup:
	eval "export GOPATH=${PWD}"

build: .setup
	export GOPATH=${PWD} && go build -v

run: .setup
	export GOPATH=${PWD} && go run ${MAIN}

clean: .setup
	go clean
