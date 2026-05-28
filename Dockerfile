# Use the official Golang Alpine image for an isolated testing environment
FROM golang:1.26-alpine

# Install necessary system tools
RUN apk add --no-cache make

# Copy the pre-built linter directly from the official image.
# This bypasses external scripts, reduces build time, and prevents checksum issues.
COPY --from=golangci/golangci-lint:latest /usr/bin/golangci-lint /usr/bin/golangci-lint

WORKDIR /app

# Cache dependencies before copying the full source code
COPY go.mod ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Run the strict linter first; if it passes, proceed with unit tests
CMD make lint && make test