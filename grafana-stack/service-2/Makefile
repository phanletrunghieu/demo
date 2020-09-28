.PHONY: build run gen-protoc

build:
	go build -o ./bin/main main.go

run:
	go run main.go

gen-protoc:
	@protoc --go_out=plugins=grpc:./proto --proto_path=proto --go_opt=paths=source_relative proto/service2.proto