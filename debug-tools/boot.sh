#!/usr/bin/env bash
if [[ $1 == *.go ]]
then
         /Users/jima/go/third-party/bin/goimports -w $1
         /usr/local/bin/gofmt -w $1
fi