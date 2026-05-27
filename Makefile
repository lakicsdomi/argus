.PHONY: test lint clean docker-test docker-down

# Makefile for Argus
# This Makefile provides commands for testing, linting,
# cleaning, and running tests in a Docker container for
# the Argus logging system.

test:
	@echo "Running tests for Argus..."
	go test ./... -v

lint:
	@echo "Formatting Go code..."
	go fmt ./...

clean:
	@echo "Cleaning up temporary log files..."
	rm -rf *.log

docker-test:
	@echo "Running tests inside an isolated Docker container..."
	docker compose up --build

docker-down:
	@echo "Stopping and removing Docker containers..."
	docker compose down