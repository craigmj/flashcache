SHELL := /bin/bash

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKG = flash
DEST = $(PKG)

bin/$(DEST): $(SRC)
	GOPATH=`pwd` go build -o bin/$(DEST) src/$(PKG)/cmd/main.go

clean:
	-@rm -rf bin/$(DEST)

prepare:
	export GOPATH=`pwd`; \
	go get -u github.com/kardianos/govendor; \
	pushd src/$(PKG); \
	../../bin/govendor init; \
	../../bin/govendor sync; \
	popd

.PHONE: clean prepare
