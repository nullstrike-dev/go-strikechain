build:
	go build -o bin/strikechain cmd/strikechain/main.go

run: build
	./bin/strikechain

test:
	go test ./...

.PHONY: build run test
