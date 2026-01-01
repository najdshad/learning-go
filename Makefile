.DEFAULT_GOAL := bin
.PHONY:fmt vet build bin clean
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
bin: build
	mv ./go-tut bin/go-tut
clean:
	rm -rf ./bin/*
