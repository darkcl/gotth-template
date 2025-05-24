package main

import (
	"gotth/server"
	"gotth/store"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[Gotth] ", log.LstdFlags)
	port := 8080
	db := store.MustNewBoltDB("gotth.db")
	defer db.Close()

	server := server.MustNewServer(logger, port, db, true)

	if err := server.Start(); err != nil {
		logger.Fatalf("Error when starting server: %s", err)
		os.Exit(1)
	}
}
