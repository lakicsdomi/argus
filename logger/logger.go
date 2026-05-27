package logger

// Defines the standard behavior for all loggers in the Argus system
type Logger interface {
	// Records a standard message belonging to a specific component
	Log(component string, message string)

	// Records a message along with an error object and its component
	LogErr(component string, message string, err error)

	// Safely releases resources (like file handles)
	Close() error
}
