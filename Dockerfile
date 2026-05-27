# Use the official Golang image for testing and linting
FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# The default command runs all unit tests with verbosity
CMD ["go", "test", "./...", "-v"]