package argus

// Logger defines the standard behavior for all loggers in the Argus system
type Logger interface {
	// Log records a standard message
	Log(message string)

	// LogErr records a message along with an error object
	LogErr(message string, err error)

	// Close safely releases resources (like file handles)
	Close() error
}
