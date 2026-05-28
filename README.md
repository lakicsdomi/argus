# 👁️ Argus Logger


Argus is a telemetry and logging library for Go with structured file logging and a built-in real-time dashboard.

The project is interface-driven, thread-safe, and designed for applications where logging and observability should stay simple and predictable.

---

## ✨ Features

### 📝 Logging

* Thread-safe structured logging
* Log levels: `VERBOSE`, `WARNING`, `ERROR`, `CRITICAL`
* Daily log file rotation
* Component-based logging (`DATABASE`, `API`, `ROUTER`, etc.)
* Interface-based design for easier testing and mocking

### 📊 Telemetry Dashboard

* Built-in HTTP dashboard
* Runs in a separate goroutine
* Real-time log viewing
* Filtering and chronological sorting

### 🚀 Development & CI

* Docker-based test and lint commands
* GitHub Actions workflow support
* Coverage badge updates using native Go tooling
* Consistent local and CI environments

---

## 📦 Installation

```bash
go get github.com/lakicsdomi/argus@v1.3.0
```

---

## 💻 Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/lakicsdomi/argus"
	"github.com/lakicsdomi/argus/dashboard"
)

func main() {
	logger, err := argus.Init("logs")
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.CloseAll()

	// Start dashboard
	go dashboard.Serve(logger.Directory, ":9090")

	logger.Verbose.Log("MAIN", "Application started")

	err = fmt.Errorf("connection timeout")
	if err != nil {
		logger.Error.LogErr("DATABASE", "operation failed", err)
	}
}
```

---

## 🧪 Development

The project uses a `Makefile` for local development and CI tasks.

| Command             | Description                                 |
| ------------------- | ------------------------------------------- |
| `make test`         | Run unit tests and generate coverage output |
| `make docker-test`  | Run tests inside Docker                     |
| `make lint`         | Format and lint locally                     |
| `make docker-lint`  | Run `golangci-lint` in Docker               |
| `make update-badge` | Update coverage badge automatically         |

---

## 🔧 Local CI Testing

You can test the GitHub Actions workflow locally with:

```bash
act
```

This simulates the GitHub Actions environment without pushing changes to GitHub.
