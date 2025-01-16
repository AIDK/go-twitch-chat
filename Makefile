all: run

build:
	@echo "Building project..."
	@go build -o bin/twitchchat

run: build
	@echo "Running the project..."
	@./bin/twitchchat