.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: start
start:
	go run ./cmd/app/main.go

.PHONY: tidy
tidy:
	go mod tidy
	go mod vendor