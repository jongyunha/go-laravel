BIN_NAME=celeritasApp

build:
	@go mod vendor
	@echo "Building Celeritas..."
	@go build -o tmp/$(BIN_NAME) .
	@echo "Done!"

run: build
	@echo "Running Celeritas..."
	@./tmp/$(BIN_NAME)
	@echo "Celeritas started!"

clean:
	@echo "Cleaning up..."
	@go clean
	@rm tmp/$(BIN_NAME)
	@echo "Done!"

test:
	@echo "Running tests..."
	@go test ./...
	@echo "Done!"

start: run