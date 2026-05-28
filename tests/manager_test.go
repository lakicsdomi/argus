package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/lakicsdomi/argus/logger"
)

func TestManager(t *testing.T) {
	dir := t.TempDir()
	m, err := logger.NewManager(dir)
	if err != nil {
		t.Fatalf("Failed to init manager: %v", err)
	}

	if m.Verbose == nil || m.Warning == nil || m.Error == nil || m.Critical == nil {
		t.Fatal("Manager loggers not fully initialized")
	}

	m.Verbose.Log("Main", "verbose msg")
	m.Warning.Log("Main", "warning msg")
	m.Error.Log("Main", "error msg")
	m.Critical.Log("Main", "critical msg")

	m.CloseAll()
}

func TestManager_InitializationError(t *testing.T) {
	tmpFile := t.TempDir() + "/blocker.log"
	_ = os.WriteFile(tmpFile, []byte("data"), 0644)

	_, err := logger.NewManager(tmpFile)
	if err == nil {
		t.Error("Expected error when manager directory is actually a file")
	}
}

// Forces failures for WARNING, ERROR, and CRITICAL loggers during initialization
func TestManager_PartialInitFailures(t *testing.T) {
	levels := []string{"WARNING", "ERROR", "CRITICAL"}

	for _, level := range levels {
		t.Run("Fail_"+level, func(t *testing.T) {
			dir := t.TempDir()

			// Create a directory exactly where the logger wants to create its file.
			// This forces os.OpenFile to fail for this specific level.
			logName := fmt.Sprintf("%s_%s.log", level, time.Now().Format("2006-01-02"))
			_ = os.Mkdir(filepath.Join(dir, logName), 0755)

			_, err := logger.NewManager(dir)
			if err == nil {
				t.Errorf("Expected error when initializing %s, got nil", level)
			}
		})
	}
}

// Ensures CloseAll handles nil loggers safely
func TestManager_CloseAllNil(t *testing.T) {
	m := &logger.Manager{}
	m.CloseAll() // Should run without panic
}
