.PHONY : all build mod run runv runb json

all: json
	@run

build: mod
	@go build

mod:
	@go mod tidy

run:
	@go run ./ --config config.json

runv:
	@go run ./ --config config.json -v

runb: build
	@./dct_backend --config config.json

json:
	easyjson ./net/structs.go

