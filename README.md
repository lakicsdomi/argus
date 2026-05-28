# 👁️ Argus Logger

![Coverage](https://img.shields.io/badge/Coverage-95.2%25-brightgreen)
![Go Version](https://img.shields.io/badge/Go-1.26-blue)
![Architecture](https://img.shields.io/badge/Architecture-SOLID-orange)
![CI/CD](https://img.shields.io/badge/CI%2FCD-Resilient-2496ED)

A robust, interface-driven telemetry and logging library for Go. Named after the all-seeing giant from Greek mythology, Argus provides thread-safe, component-segregated file logging alongside a concurrent, real-time HTTP telemetry dashboard.

Engineered with **Clean Architecture** and **DevOps best practices** in mind, Argus is designed to be easily mockable, highly observable, and CI/CD-ready.

---

## ✨ Key Features

### 🏗 Software Architecture
- **Interface-Driven (SOLID)**: Core logging mechanisms are abstracted behind interfaces, eliminating global state and allowing for seamless dependency injection and mocking in your applications.
- **Concurrent Telemetry**: Features an out-of-the-box, standalone HTTP dashboard running in a dedicated `goroutine` to view, filter, and sort logs chronologically without blocking the main application thread.
- **Component Segregation**: Enforces a mandatory `COMPONENT` identifier (e.g., `DATABASE`, `ROUTER`) to ensure highly organized, traceable, and grep-friendly structured logs.
- **Multi-Level Verbosity**: Native support for `VERBOSE`, `WARNING`, `ERROR`, and `CRITICAL` levels, automatically rolling into separate daily log files.

### 🚀 DevOps & Infrastructure
- **Deterministic CI Environments**: Incorporates Docker-isolated static analysis (`golangci-lint`) and testing targets to eliminate "Works on my machine" anomalies.
- **Self-Healing CI Pipeline**: GitHub Actions workflow features a fallback mechanism—if the native Action linter fails due to environment mismatch, it automatically falls back to a containerized Linter.
- **Automated Coverage Badges**: Zero reliance on third-party opaque badge generators. Coverage is dynamically extracted via native Go tools and updated via `sed` scripting in the `Makefile`.

---

## 📦 Installation

```bash
go get github.com/lakicsdomi/argus@v1.3.0

```

---

## 💻 Usage Example

Integrating Argus into your application is straightforward and enforces clean boundaries:

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
	// Ensure file descriptors are gracefully closed on shutdown
	defer logger.CloseAll()

	// Spin up the real-time telemetry dashboard in a non-blocking goroutine
	go dashboard.Serve(logger.Directory, ":9090")

	// Standard structured logging
	logger.Verbose.Log("MAIN", "Application initialized successfully.")

	// Error logging with context
	err = fmt.Errorf("connection timeout")
	if err != nil {
		logger.Error.LogErr("DATABASE", "An operation failed", err)
	}
}

```

---

## 🧪 Development & CI/CD

Argus relies on a robust `Makefile` as the single source of truth for both local development and CI/CD pipelines.

### Makefile Commands

| Command | Description |
| --- | --- |
| `make test` | Runs local unit tests and generates a `coverage.out` profile. |
| `make docker-test` | **[CI Standard]** Runs tests inside an isolated Docker container ensuring 100% environmental parity, extracting coverage dynamically. |
| `make lint` | Formats and lints the codebase locally. |
| `make docker-lint` | Runs strict static analysis (`golangci-lint`) inside a multi-stage Docker container. |
| `make update-badge` | Parses the coverage profile and updates this README's badge automatically. |

### CI/CD Simulation

To validate the `.github/workflows/ci.yml` pipeline locally without pushing to GitHub, this repository supports [act](https://github.com/nektos/act). Run `act` in the root directory to simulate the exact GitHub Actions runner environment safely on your machine.
