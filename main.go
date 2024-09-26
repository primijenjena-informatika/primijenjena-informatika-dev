package main

import (
	"embed"
	"log"
	"primijenjena-informatika-dev/cmd/server"

	"github.com/joho/godotenv"
)

//go:embed public/*
var publicFS embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.SetupConfig()
	srv := server.CreateServer(publicFS)
	server.StartServer(srv)
}
