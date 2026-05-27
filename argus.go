package argus

import "fmt"

// Argus holds multiple loggers based on verbosity levels
type Argus struct {
	Verbose  Logger
	Warning  Logger
	Error    Logger
	Critical Logger
}

// Init initializes the complete Argus logging system
func Init(directory string) (*Argus, error) {
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

	return &Argus{
		Verbose:  verbose,
		Warning:  warning,
		Error:    errLogger,
		Critical: critical,
	}, nil
}

// CloseAll securely closes all active log files
func (a *Argus) CloseAll() {
	if a.Verbose != nil {
		_ = a.Verbose.Close()
	}
	if a.Warning != nil {
		_ = a.Warning.Close()
	}
	if a.Error != nil {
		_ = a.Error.Close()
	}
	if a.Critical != nil {
		_ = a.Critical.Close()
	}
}
