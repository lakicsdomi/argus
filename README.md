# Argus Logger

Argus is a lightweight, multi-leveled logging library for Go. Named after the all-seeing giant from Greek mythology, it provides clear, segregated logging into specific files based on verbosity levels and application components.

## Features
- **Multi-Level Verbosity**: Supports `VERBOSE`, `WARNING`, `ERROR`, and `CRITICAL` log levels.
- **Structured Logging**: Requires a mandatory component identifier (e.g., `DATABASE`, `ROUTER`) to ensure highly organized and traceable logs.
- **File Segregation**: Automatically creates separate daily log files for each verbosity level.
- **Telemetry Dashboard**: An out-of-the-box, standalone HTTP dashboard to view, filter, and sort logs chronologically in a modern UI.
- **Interface-Driven**: Built on SOLID principles, allowing for easy mocking and extending.

## Installation

```bash
go get github.com/lakicsdomi/argus@v1.3.0
```

## Usage example
```go
package main

import (
    "fmt"
    "log"
    "github.com/lakicsdomi/argus"
    "github.com/lakicsdomi/argus/dashboard"
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