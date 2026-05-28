package tests

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/lakicsdomi/argus/logger"
)

// Tests the basic functionality of FileLogger, including logging and error handling
func TestFileLogger(t *testing.T) {
	dir := t.TempDir()
	l, err := logger.NewFileLogger(dir, "TEST")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}
	defer func() {
		if err := l.Close(); err != nil {
			t.Errorf("Failed to close logger: %v", err)
		}
	}()

	l.Log("TestComp", "standard message")
	l.LogErr("TestComp", "error message", errors.New("test error"))
	l.LogErr("TestComp", "nil error", nil)
}

// Tests that NewFileLogger returns an error when given an invalid directory path
func TestFileLogger_InvalidPath(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "blocker")
	_ = os.WriteFile(tmpFile, []byte("i am a file"), 0644)

	_, err := logger.NewFileLogger(tmpFile, "FAIL")
	if err == nil {
		t.Error("Expected error when using a file path as a log directory, got nil")
	}
}

// Tests that Close() handles a nil file gracefully (for coverage)
func TestFileLogger_CloseNil(t *testing.T) {
	l, _ := logger.NewFileLogger(t.TempDir(), "DEBUG")
	if err := l.Close(); err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}
