package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lakicsdomi/argus"
)

// Tests the initialization of the Argus logging system and verifies that log files are created
func TestArgusInitialization(t *testing.T) {
	tempDir := t.TempDir()

	manager, err := argus.Init(tempDir)
	if err != nil {
		t.Fatalf("Expected no error during initialization, got %v", err)
	}
	if manager == nil {
		t.Fatal("Expected a valid Manager instance, got nil")
	}

	manager.CloseAll()

	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read temp directory: %v", err)
	}
	if len(files) == 0 {
		t.Error("Expected log files to be created, but directory is empty")
	}
}

// Tests that initializing Argus with a file path instead of a directory returns an error
func TestArgusInitialization_InvalidDirectory(t *testing.T) {

	tmpFile := filepath.Join(t.TempDir(), "blocked.log")
	if err := os.WriteFile(tmpFile, []byte("I am a file, not a directory"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := argus.Init(tmpFile)

	if err == nil {
		t.Error("Expected an error when initializing with a file path instead of a directory, got nil")
	}
}
