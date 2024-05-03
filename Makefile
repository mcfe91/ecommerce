build:
	@go build -o bin/ecommerce

run: build
	@./bin/ecommerce

test:
	@go test -v ./...