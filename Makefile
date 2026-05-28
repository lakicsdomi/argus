.PHONY: deps lint-simple test test-no-coverage lint clean docker-test docker-down coverage-summary update-badge docker-lint 

# Makefile for Argus
# Provides commands for dependency management, testing, linting,
# cleaning, and isolated execution via Docker.

deps:
	@echo "Downloading and tidying dependencies..."
	go mod download
	go mod tidy

lint-simple:
	@echo "Running simple lint (vet, fmt)..."
	go vet ./...
	go fmt ./...

test:
	@echo "Running tests with coverage for all packages..."
	go test -v ./... -coverpkg=./... -coverprofile=coverage.out
	@echo "Report coverage..."
	go tool cover -func=coverage.out

coverage-summary:
	@echo "Extracting total coverage..."
	@go tool cover -func=coverage.out | grep total | awk '{print $$3}'

# Dynamically updates the coverage badge in README.md without relying on 3rd party actions
update-badge:
	@echo "Updating coverage badge in README.md..."
	@COVERAGE=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}' | tr -d '%'); \
	sed -i -e "s/Coverage-[0-9.]*%25/Coverage-$${COVERAGE}%25/g" README.md

test-no-coverage:
	@echo "Running tests without coverage for faster execution..."
	go test ./... -v

lint:
	@echo "Running strict static analysis (golangci-lint)..."
	golangci-lint run ./...

clean:
	@echo "Cleaning up temporary log files and coverage reports..."
	rm -rf *.log
	rm -f coverage.out

docker-test:
	@echo "Running tests inside an isolated Docker container..."
	docker compose up --build --abort-on-container-exit --exit-code-from test
# --abort-on-container-exit ensures that if the test container exits (either success or failure), the entire compose session will stop, allowing us to capture the exit code properly.
# --exit-code-from test ensures that the exit code of the test container is propagated to the docker compose command, so we can detect test failures in CI.

docker-down:
	@echo "Stopping and removing Docker containers..."
	docker compose down

# This target allows running golangci-lint in an isolated Docker container, 
# ensuring a consistent linting environment regardless of the host setup. 
# Pretty useful for ACT local testing of the CI pipeline.
docker-lint:
	@echo "Running golangci-lint in an isolated Docker container..."
	docker compose run --rm --build test make lint