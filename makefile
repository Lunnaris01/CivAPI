.PHONY: run build

run:
	go run ./cmd/CivAPI

build:
	go build -o civapi ./cmd/CivAPI
