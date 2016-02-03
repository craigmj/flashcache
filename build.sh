#!/bin/bash
set -e
export GOPATH=`pwd`
for p in 'github.com/craigmj/commander' \
	'github.com/golang/glog' \
	; do
	if [ ! -d src/$p ]; then
		go get $p
	fi
done
if [ ! -d bin ]; then
	mkdir bin
fi
go build -o bin/flashcache src/cmd/flashcache.go