M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: setup dep schema clean build local-db run

setup: 
	$(info $(M) Fetching github.com/golang/dep...)
	go get github.com/golang/dep/cmd/dep
	$(info $(M) Fetching github.com/jteeuwen/go-bindata/...)
	go get -u github.com/jteeuwen/go-bindata/...

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...)
	dep ensure -v

schema: ; $(info $(M) Embedding schema files into binary...)
	go-bindata -ignore=\.go -pkg=schema -o=./schema/bindata.go schema/*
	
clean: ; $(info $(M) Removing generated files... )
	$(RM) schema/bindata.go

build: schema ; $(info $(M) Building project...)
	go build -o main ./cmd/main.go

local-db:
	@docker-compose down
	@docker-compose up -d

run: build
	PORT=3000 ./main