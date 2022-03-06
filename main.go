package main

import (
	server "bareksa-interview-project/interfaces/http"
	log "bareksa-interview-project/util/logger"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[!] Please setup a .env file first")
		os.Exit(1)
	}

	// Get LOG_FILE from .env
	filename := fmt.Sprintf("%s.txt", os.Getenv("LOG_FILE"))
	if filename == ".txt" {
		fmt.Println("[!] Log file should not be empty (make sure LOG_FILE property in .env is set correctly)")
		os.Exit(1)
	}

	// Setup log file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("[!] Error creating log file: %s\n", err)
		os.Exit(1)
	}

	// Initialize logger
	customLogger := log.CreateLogger(file)

	defer func() {
		if r := recover(); r != nil {
			customLogger.WriteLog(log.ERROR, r)
		}
	}()

	isDebugMode, err := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	if err != nil {
		fmt.Println("[!] DEBUG_MODE property in .env is not set correctly")
		os.Exit(1)
	}

	dbPass := os.Getenv("POSTGRES_PASSWORD")
	if dbPass == "" {
		fmt.Println("[!] POSTGRES_PASSWORD property in .env is not set correctly")
		os.Exit(1)
	}

	migrationPass := os.Getenv("MIGRATION_PASSWORD")
	if migrationPass == "" {
		fmt.Println("[!] MIGRATION_PASSWORD property in .env is not set correctly")
		os.Exit(1)
	}

	server.Start(isDebugMode, dbPass, migrationPass, &customLogger)
}
