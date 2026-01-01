.DEFAULT_GOAL := bin
.PHONY:fmt vet build bin clean
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build
bin: build
	[ -d ./bin ] || mkdir -p ./bin && mv ./learning-go bin/learning-go
clean:
	rm -rf ./bin/*
run:
	./bin/learning-go
