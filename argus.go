package argus

import (
	"github.com/lakicsdomi/argus/logger"
)

// Initializes the complete Argus logging system
func Init(directory string) (*logger.Manager, error) {
	return logger.NewManager(directory)
}
