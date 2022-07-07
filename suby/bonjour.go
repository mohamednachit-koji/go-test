package suby

import (
	"go.uber.org/zap"
)

// Function that says bonjour
func Bonjour() {
	logger, _ := zap.NewProduction()
	logger.Info("Bonjour monsieur")
}
