package argus

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// FileLogger implements the Logger interface and writes to a specific file
type FileLogger struct {
	levelName string
	file      *os.File
}

// NewFileLogger creates a new FileLogger instance for a specific severity level
func NewFileLogger(directory, levelName string) (*FileLogger, error) {
	if err := os.MkdirAll(directory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	// Generate the file name based on the level and current date
	fileName := fmt.Sprintf("%s_%s.log", levelName, time.Now().Format("2006-01-02"))
	filePath := filepath.Join(directory, fileName)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s log file: %w", levelName, err)
	}

	return &FileLogger{
		levelName: levelName,
		file:      file,
	}, nil
}

// Log writes a standard message to both the console and the specific log file
func (l *FileLogger) Log(message string) {
	l.write(fmt.Sprintf("%s", message))
}

// LogErr writes an error message to both the console and the specific log file
func (l *FileLogger) LogErr(message string, err error) {
	if err == nil {
		return
	}
	l.write(fmt.Sprintf("%s - Error details: %v", message, err))
}

// write handles the actual formatting and I/O operations
func (l *FileLogger) write(content string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] %s: %s\n", timestamp, l.levelName, content)

	// Print to standard output (console)
	log.Print(logMessage)

	// Write the message to the file
	if l.file != nil {
		if _, writeErr := l.file.WriteString(logMessage); writeErr != nil {
			log.Printf("Failed to write to %s log file: %v\n", l.levelName, writeErr)
		}
	}
}

// Close safely closes the underlying log file
func (l *FileLogger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
