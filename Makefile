
build:
	go build

test: build
	@PATH=${PWD}/test:${PATH} ./test/*.sheep
