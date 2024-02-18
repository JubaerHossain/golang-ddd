install:
	go mod tidy
dev:
	go run ./cmd/server/server.go

build:
	go build -o ./bin/server ./cmd/server/server.go

run:
	./bin/server