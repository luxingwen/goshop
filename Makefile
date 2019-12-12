export GOPROXY=https://goproxy.io
default: build

build: export GO111MODULE=on

build:
	go build -o restful/bin/goshop-restful restful/main.go


build-swag:
	cd restful && swag init
	go build -o restful/bin/goshop-restful restful/main.go


build-swag-run: build-swag
	cd restful && ./bin/goshop-restful


test:
	go build -o restful/bin/goshop-test restful/test/main.go
	restful/bin/goshop-test