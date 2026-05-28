package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Implements the Logger interface and writes to a specific file
type FileLogger struct {
	levelName string
	file      *os.File
}

// Creates a new FileLogger instance for a specific severity level
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

// Writes a standard message to both the console and the specific log file
func (l *FileLogger) Log(component string, message string) {
	l.write(component, message)
}

// Writes an error message to both the console and the specific log file
func (l *FileLogger) LogErr(component string, message string, err error) {
	if err == nil {
		return
	}
	l.write(component, fmt.Sprintf("%s - Error details: %v", message, err))
}

// Handles the actual formatting and I/O operations
func (l *FileLogger) write(component string, content string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// [TIMESTAMP] LEVEL: COMPONENT: MESSAGE
	logMessage := fmt.Sprintf("[%s] %s: %s: %s\n", timestamp, l.levelName, component, content)

	// Only output to the console if the level is CRITICAL or ERROR
	if l.levelName == "CRITICAL" || l.levelName == "ERROR" {
		log.Print(logMessage)
	}

	if l.file != nil {
		if _, writeErr := l.file.WriteString(logMessage); writeErr != nil {
			log.Printf("failed to write to %s log file: %v\n", l.levelName, writeErr)
		}
	}
}

// Safely closes the underlying log file
func (l *FileLogger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
