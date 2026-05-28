package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/lakicsdomi/argus/dashboard"
)

// Tests the dashboard's ability to serve log entries and handle various edge cases
func TestDashboard(t *testing.T) {
	tmp := t.TempDir()

	// valid log file
	validLog := filepath.Join(tmp, "VERBOSE_2026-05-28.log")
	err := os.WriteFile(validLog, []byte("[2026-05-28 10:00:00] VERBOSE: Comp: Msg\n"), 0644)
	if err != nil {
		t.Fatalf("failed to create valid log: %v", err)
	}

	// should be ignored
	err = os.WriteFile(filepath.Join(tmp, "notes.txt"), []byte("ignore"), 0644)
	if err != nil {
		t.Fatalf("failed to create ignored file: %v", err)
	}

	// fake .log directory
	err = os.Mkdir(filepath.Join(tmp, "WARNING_2026-05-28.log"), 0755)
	if err != nil {
		t.Fatalf("failed to create fake log dir: %v", err)
	}

	// unreadable file
	unreadable := filepath.Join(tmp, "ERROR_2026-05-28.log")
	err = os.WriteFile(unreadable, []byte("hidden"), 0222)
	if err != nil {
		t.Fatalf("failed to create unreadable log: %v", err)
	}

	port := ":8181"

	go dashboard.Serve(tmp, port)

	time.Sleep(200 * time.Millisecond)

	t.Run("root endpoint", func(t *testing.T) {
		resp, err := http.Get("http://localhost" + port + "/")
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("logs endpoint", func(t *testing.T) {
		resp, err := http.Get("http://localhost" + port + "/api/logs")
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		var logs []dashboard.LogEntry

		if err := json.NewDecoder(resp.Body).Decode(&logs); err != nil {
			t.Fatalf("decode failed: %v", err)
		}

		if len(logs) == 0 {
			t.Fatal("expected at least one parsed log entry")
		}
	})
}

// Tests that the dashboard handles an invalid log directory gracefully
func TestDashboard_InvalidDirectory(t *testing.T) {
	port := ":8182"

	go dashboard.Serve("/path/that/does/not/exist", port)

	time.Sleep(200 * time.Millisecond)

	resp, err := http.Get("http://localhost" + port + "/api/logs")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed reading body: %v", err)
	}

	if len(body) == 0 {
		t.Fatal("expected some kind of response body")
	}
}

// Tests that the dashboard handles an invalid port format gracefully
func TestDashboard_ServePortError(t *testing.T) {
	dashboard.Serve(t.TempDir(), ":-1")
}
