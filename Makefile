OUTPUT=bt-get
DIR=./cmd/bt-get

all: build

build:
	go build -o ${OUTPUT} ${DIR}