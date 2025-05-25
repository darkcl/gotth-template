package main

import (
	"gotth/server"
	"gotth/store"
	"log"
	"os"
	"strconv"
)

func getPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if portInt, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid port: %s", port)
		os.Exit(1)
	} else {
		return portInt
	}

	return 8080
}

func getUseEmbedded() bool {
	isEmbedded := os.Getenv("USE_EMBEDDED")
	if isEmbedded == "" {
		return false
	}

	if isEmbeddedBool, err := strconv.ParseBool(isEmbedded); err != nil {
		log.Fatalf("Invalid embedded flag: %s", isEmbedded)
		os.Exit(1)
	} else {
		return isEmbeddedBool
	}

	return false
}

func main() {
	logger := log.New(os.Stdout, "[Gotth] ", log.LstdFlags)
	port := getPort()
	shouldUseEmbedded := getUseEmbedded()

	db := store.MustNewBoltDB("gotth.db")
	defer db.Close()

	server := server.MustNewServer(logger, port, db, shouldUseEmbedded)

	if err := server.Start(); err != nil {
		logger.Fatalf("Error when starting server: %s", err)
		os.Exit(1)
	}
}
