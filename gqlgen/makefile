M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: setup dep schema clean build local-db run

setup: 
	$(info $(M) Fetching github.com/golang/dep...)
	go get github.com/golang/dep/cmd/dep
	$(info $(M) Fetching github.com/jteeuwen/go-bindata/...)
	go get -u github.com/jteeuwen/go-bindata/...

dep: setup ; $(info $(M) Ensuring vendored dependencies are up-to-date...)
	dep ensure -v

prisma-gen: ; $(info $(M) Generating prisma files...)
	prisma generate

prisma-deploy: ; $(info $(M) Deploy prisma...)
	prisma deploy --env-file .env
	
gqlgen: ; $(info $(M) Generating gqlgen files...)
	go run scripts/gqlgen.go -v

build: schema ; $(info $(M) Building project...)
	go build -o main ./cmd/main.go

local-db:
	@docker-compose down
	@docker-compose up -d

run: build
	PORT=3000 ./main
