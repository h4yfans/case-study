package environment

import (
	"os"
	"strings"
)

const (
	DefaultLogLevel = "ERROR"
)

func LogLevel() string {
	if level := strings.ToUpper(os.Getenv("LOG_LEVEL")); level != "" {
		return level
	}
	return DefaultLogLevel
}
