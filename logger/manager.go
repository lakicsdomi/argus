package logger

import "fmt"

// Holds multiple loggers based on verbosity levels and the target directory
type Manager struct {
	Directory string
	Verbose   Logger
	Warning   Logger
	Error     Logger
	Critical  Logger
}

// Initializes the complete logging system
func NewManager(directory string) (*Manager, error) {
	verbose, err := NewFileLogger(directory, "VERBOSE")
	if err != nil {
		return nil, fmt.Errorf("verbose logger init failed: %w", err)
	}

	warning, err := NewFileLogger(directory, "WARNING")
	if err != nil {
		return nil, fmt.Errorf("warning logger init failed: %w", err)
	}

	errLogger, err := NewFileLogger(directory, "ERROR")
	if err != nil {
		return nil, fmt.Errorf("error logger init failed: %w", err)
	}

	critical, err := NewFileLogger(directory, "CRITICAL")
	if err != nil {
		return nil, fmt.Errorf("critical logger init failed: %w", err)
	}

	return &Manager{
		Directory: directory,
		Verbose:   verbose,
		Warning:   warning,
		Error:     errLogger,
		Critical:  critical,
	}, nil
}

// Securely closes all active log files
func (m *Manager) CloseAll() {
	if m.Verbose != nil {
		_ = m.Verbose.Close()
	}
	if m.Warning != nil {
		_ = m.Warning.Close()
	}
	if m.Error != nil {
		_ = m.Error.Close()
	}
	if m.Critical != nil {
		_ = m.Critical.Close()
	}
}
