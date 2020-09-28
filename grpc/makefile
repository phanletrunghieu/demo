.PHONY: gen-protoc server client

PROJECT_PATH = $(shell pwd)

gen-protoc:
	@protoc -I $(GOPATH)/src $(PROJECT_PATH)/proto/domain/domain.proto --go_out=plugins=grpc:$(GOPATH)/src
	@protoc -I $(GOPATH)/src $(PROJECT_PATH)/proto/service/service.proto --go_out=plugins=grpc:$(GOPATH)/src

server:
	@go run server/main.go

client:
	@go run client/main.go