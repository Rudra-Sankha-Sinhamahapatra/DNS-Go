package utils_test

import (
	"example/user/hello/src/utils"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Changing the working directory to the project root
	err := os.Chdir("..") // Move up one directory
	if err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}

	config, err := utils.LoadConfig()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	if config.Ip != "127.0.0.1" && config.Ip != "0.0.0.0" {
		t.Errorf("Expected Ip 127.0.0.1 or 0.0.0.0, got %s", config.Ip)
	}

	if config.ServerPort != 9000 {
		t.Errorf("Expected port 9000, got %d", config.ServerPort)
	}
	if config.LogFile != "dns_server.log" {
		t.Errorf("Expected log file dns_server.log, got %s", config.LogFile)
	}
}
