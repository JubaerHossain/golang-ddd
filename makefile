.PHONY: proto
proto:
	@./scripts/proto.sh trainer
dev:
	go run ./cmd/server/server.go