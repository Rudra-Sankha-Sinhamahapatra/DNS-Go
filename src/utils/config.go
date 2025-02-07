package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	LogFile    string
}

var AppConfig Config

func LoadConfig() {
	_ = godotenv.Load()

	portStr := os.Getenv("DNS_SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 8090
	}

	logFile := os.Getenv("DNS_LOG_FILE")
	if logFile == "" {
		logFile = "dns_server.log"
	}

	AppConfig = Config{
		ServerPort: port,
		LogFile:    logFile,
	}

}
