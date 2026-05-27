package dashboard

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Represents a single parsed line from a log file
type LogEntry struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// Starts an independent HTTP server for the Argus dashboard
// It runs on the specified port and reads logs from the given directory
func Serve(logDirectory string, port string) {
	mux := http.NewServeMux()

	// Serve the embedded HTML UI
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(dashboardHTML))
	})

	// API endpoint to fetch the latest logs as JSON
	mux.HandleFunc("/api/logs", func(w http.ResponseWriter, r *http.Request) {
		logs := readLogs(logDirectory)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(logs)
	})

	log.Printf("Argus Telemetry Dashboard running on http://localhost%s\n", port)

	// Start the standalone server
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Printf("Argus Dashboard server stopped: %v", err)
	}
}

// readLogs scans the directory and parses the log files into structured data
func readLogs(directory string) []LogEntry {
	var entries []LogEntry

	files, err := os.ReadDir(directory)
	if err != nil {
		return entries
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".log") {
			continue
		}

		content, err := os.ReadFile(filepath.Join(directory, file.Name()))
		if err != nil {
			continue
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}

			// Simple parsing based on the file name prefix (e.g., VERBOSE_2026-05-27.log)
			level := strings.Split(file.Name(), "_")[0]
			entries = append(entries, LogEntry{
				Level:   level,
				Message: line,
			})
		}
	}

	return entries
}
