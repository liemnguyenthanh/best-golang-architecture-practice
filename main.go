// cmd/main.go
package main

import (
	"api-instagram/config"
	"api-instagram/db"
	"api-instagram/internal/server"
	"log"
)

func main() {
	cfg := config.NewConfig()
	log.Println("Configuration PORT", cfg)
	db := db.GetMySQLInstance(cfg, true)
	s := server.NewServer(cfg, db)

	s.Run()
}
