PWD := $(shell pwd)
BIN := $(PWD)/bin

VERSION := $(shell git rev-list HEAD | head -1)

BUILD_DATE := $(shell date +%Y-%m-%d\ %H:%M:%S)

default: test

test:
	echo $(PWD)
	echo $(BIN)
	echo $(VERSION)
	echo $(BUILD_DATE)
	