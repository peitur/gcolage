GOPATH=${PWD}
GCOLAGE=gcolage.go
FTREECHKS=ftreechksum.go

GCOLAGEBIN=gcolage
FTREECHKSBIN=ftreechksum

all: .setup gcolage ftreechksum

.setup:
	eval "export GOPATH=${PWD}"

gcolage: .setup
	export GOPATH=${PWD} && go build -v -o ${GCOLAGEBIN} ${GCLAGE}

ftreechksum: .setup
	export GOPATH=${PWD} && go build -v -o ${FTREECHKSBIN} ${FTREECHKS}


clean: .setup
	go clean
