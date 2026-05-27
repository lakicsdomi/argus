package manager_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lakicsdomi/argus/logger"
)

// Tests the initialization of the Manager and its loggers
func TestNewManager(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	levels := []string{"VERBOSE", "WARNING", "ERROR", "CRITICAL"}

	// Act
	manager, err := logger.NewManager(dir)
	defer manager.CloseAll()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if manager == nil {
		t.Fatal("Expected a valid Manager instance, got nil")
	}

	// Verify that all requested log files were actually created
	for _, level := range levels {
		matches, matchErr := filepath.Glob(filepath.Join(dir, level+"_*.log"))
		if matchErr != nil || len(matches) == 0 {
			t.Errorf("Expected log file for level %s to be created, but it was not", level)
		}
	}
}

// Tests if the FileLogger correctly formats and writes data to the file
func TestFileLogger_Log(t *testing.T) {
	// Arrange
	dir := t.TempDir()
	testComponent := "DATABASE"
	testMessage := "Connection successful."

	// Creating the logger is part of the test setup (Arrange)
	fileLogger, err := logger.NewFileLogger(dir, "TEST")
	if err != nil {
		t.Fatalf("Failed to create file logger, got %v", err)
	}

	// Act
	fileLogger.Log(testComponent, testMessage)
	fileLogger.Close() // Ensure buffers are flushed to the file before reading

	// Assert
	matches, _ := filepath.Glob(filepath.Join(dir, "TEST_*.log"))
	if len(matches) == 0 {
		t.Fatal("Log file was not created")
	}

	content, readErr := os.ReadFile(matches[0])
	if readErr != nil {
		t.Fatalf("Failed to read log file: %v", readErr)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, testComponent) {
		t.Errorf("Expected log to contain component %s", testComponent)
	}
	if !strings.Contains(contentStr, testMessage) {
		t.Errorf("Expected log to contain message %s", testMessage)
	}
}
