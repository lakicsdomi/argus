![Coverage](https://raw.githubusercontent.com/lakicsdomi/argus/main/coverage.svg)

# Argus Logger

Argus is a lightweight, multi-leveled logging library for Go. Named after the all-seeing giant from Greek mythology, it provides clear, segregated logging into specific files based on verbosity levels and application components.

## Features
- **Multi-Level Verbosity**: Supports `VERBOSE`, `WARNING`, `ERROR`, and `CRITICAL` log levels.
- **Structured Logging**: Requires a mandatory component identifier (e.g., `DATABASE`, `ROUTER`) to ensure highly organized and traceable logs.
- **File Segregation**: Automatically creates separate daily log files for each verbosity level.
- **Telemetry Dashboard**: An out-of-the-box, standalone HTTP dashboard to view, filter, and sort logs chronologically in a modern UI (utilizing Go `goroutines`).
- **Interface-Driven**: Built on SOLID principles, allowing for easy mocking and extending without global states.

## Installation

```bash
go get [github.com/lakicsdomi/argus@v1.3.0](https://github.com/lakicsdomi/argus@v1.3.0)
```

## Usage example

```go
package main

import (
    "fmt"
    "log"
    "[github.com/lakicsdomi/argus](https://github.com/lakicsdomi/argus)"
    "[github.com/lakicsdomi/argus/dashboard](https://github.com/lakicsdomi/argus/dashboard)"
)

func main() {
    // Initialize the Argus logging manager
    logger, err := argus.Init("logs")
    if err != nil {
        log.Fatalf("Failed to initialize logger: %v", err)
    }
    defer logger.CloseAll()

    // Start the built-in telemetry dashboard in a separate goroutine
    go dashboard.Serve(logger.Directory, ":9090")

    // Log a standard message with a mandatory COMPONENT identifier
    logger.Verbose.Log("MAIN", "Application initialized successfully.")

    // Log an error
    err = fmt.Errorf("connection timeout")
    if err != nil {
        logger.Error.LogErr("DATABASE", "An operation failed", err)
    }
}
```

## Development and testing
Argus includes a `Makefile` to simplify the development workflow.
- Run standard unit tests: `make test`
- Format the codebase: `make lint`
- Run tests in a completely isolated Docker environment (No local Go installation required!): `make docker-test`

## CI/CD
This project uses **GitHub Actions** to enforce code quality and run tests on every push and pull request.
To validate the CI pipeline locally without pushing to GitHub, I highly recommend using [act](https://github.com/nektos/act). Install `act` and run it in the terminal to simulate the CI environment locally.
