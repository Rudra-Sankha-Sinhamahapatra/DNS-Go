package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	LogFile    string
	Ip         string
}

var AppConfig Config

func LoadConfig() (Config, error) {
	err := godotenv.Load()

	if err != nil {
		return Config{}, err
	}

	portStr := os.Getenv("DNS_SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 8900
	}

	logFile := os.Getenv("DNS_LOG_FILE")
	if logFile == "" {
		logFile = "dns_server.log"
	}

	ip := os.Getenv("DNS_SERVER_IP")
	if ip == "" {
		ip = "0.0.0.0" // Default to all network interfaces
	}

	AppConfig = Config{
		ServerPort: port,
		LogFile:    logFile,
		Ip:         ip,
	}

	return AppConfig, nil

}
