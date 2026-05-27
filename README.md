# Argus Logger

Argus is a lightweight, multi-leveled logging library for Go. Named after the all-seeing giant from Greek mythology, it provides clear, segregated logging into specific files based on verbosity levels.

## Features
- **Multi-Level Verbosity**: Supports `VERBOSE`, `WARNING`, `ERROR`, and `CRITICAL` log levels.
- **File Segregation**: Automatically creates separate daily log files for each verbosity level.
- **Interface-Driven**: Built on SOLID principles, allowing for easy mocking and extending (e.g., adding network loggers).

## Installation

```bash
go get [github.com/lakicsdomi/argus@v1.0.0](https://github.com/lakicsdomi/argus@v1.0.0)
```

## Usage example
```go
package main

import (
	"log"
	"[github.com/lakicsdomi/argus](https://github.com/lakicsdomi/argus)"
)

func main() {
	// Initialize the Argus logging manager
	logger, err := argus.Init("logs")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.CloseAll()

	// Log a standard message
	logger.Verbose.Log("Application initialized successfully.")

	// Log an error
	err = doSomething()
	if err != nil {
		logger.Error.LogErr("An operation failed", err)
	}
}
```

## Upcoming Features
- **Web Dasboard**: An out-of-the-box telemetry endpoint to view logs in a UI (similar to Prometheus/Grafana concepts).