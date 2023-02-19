BINARY_NAME=dotfair

dependencies:
	@echo "Installing dependencies..."
	@go mod download

test: dependencies
	@echo "Running tests..."
	@go test -v ./...

e2e-test: dependencies
	@echo "Running end-to-end tests..."
	@echo "Not implemented yet"


build: dependencies
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) -v

run: build
	@echo "Running..."
	@./bin/$(BINARY_NAME)

docker-build: dependencies
	@echo "Building docker image..."
	@docker build -t dotfair .

clean:
	@echo "Cleaning..."
	@rm -rf bin/

.PHONY: dependencies test build run clean
