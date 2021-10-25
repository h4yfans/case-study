package environment

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	DefaultContextTimeout = time.Minute * 5 // 5 Minute
	DefaultPort           = 8080
)

func Debug() bool {
	return strings.ToUpper(os.Getenv("DEBUG")) == "TRUE"
}

func BoilDebug() bool {
	return strings.ToUpper(os.Getenv("BOIL_DEBUG")) == "TRUE"
}

func ContextTimeout() time.Duration {
	env := os.Getenv("CONTEXT_TIMEOUT")
	if env == "" {
		return DefaultContextTimeout
	}

	timeout, err := strconv.Atoi(env)
	if err != nil {
		zap.L().Fatal("Context timeout env could not cast to int", zap.Error(err), zap.String("env", env))
	}
	return time.Duration(timeout) * time.Second
}

func Port() int {
	env := os.Getenv("PORT")
	if env == "" {
		return DefaultPort
	}
	port, err := strconv.Atoi(env)
	if err != nil {
		zap.L().Fatal("Port env could not cast to int", zap.Error(err), zap.String("env", env))
	}
	return port
}

// ReadEnvFile is for only local development
func ReadEnvFile() {
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			errParent := godotenv.Load("../.env")
			if errParent != nil {
				zap.L().Error("Cannot read ../.env file", zap.Error(errParent))
			}
		}
	}
}
