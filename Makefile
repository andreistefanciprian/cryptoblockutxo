
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

proto:
	@protoc \
	--go_out=paths=source_relative:. \
	--go-grpc_out=paths=source_relative:. \
	./proto/*.proto

.PHONY: test test-coverage proto