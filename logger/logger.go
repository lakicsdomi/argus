package logger

// Defines the standard behavior for all loggers in the Argus system
type Logger interface {
	// Records a standard message
	Log(message string)

	// Records a message along with an error object
	LogErr(message string, err error)

	// Safely releases resources (like file handles)
	Close() error
}
