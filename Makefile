export GOPROXY=https://goproxy.io
default: build

build: export GO111MODULE=on

build:
	go build -o restful/bin/goshop-restful restful/main.go
