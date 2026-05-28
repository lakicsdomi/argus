package logger

import (
	"testing"
)

// Tests internal methods with a nil file pointer
func TestFileLogger_NilFileCoverage(t *testing.T) {
	l := &FileLogger{
		levelName: "DEBUG",
		file:      nil,
	}

	// Check that write still works without a file and doesn't panic
	l.Log("Test", "msg")

	// checks that it doesn't panic and returns nil
	if err := l.Close(); err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

// Tests internal write failure when the file handle is closed unexpectedly
func TestFileLogger_WriteErrorCoverage(t *testing.T) {
	dir := t.TempDir()
	l, err := NewFileLogger(dir, "DEBUG")
	if err != nil {
		t.Fatalf("Setup failed: %v", err)
	}

	// Force close the underlying file to trigger a WriteString error
	_ = l.file.Close()

	// This hits the writeErr != nil branch inside write()
	l.Log("Test", "This should fail to write")
}
