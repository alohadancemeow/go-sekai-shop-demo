package main

import (
	"context"
	"log"
	"os"

	"github.com/Alohadancemeow/go-sekai-shop-demo/config"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	log.Println(cfg)

	// Database connection
	// db := database.DbConn(ctx, &cfg)
	// defer db.Disconnect(ctx)

	// // Start Server
	// server.Start(ctx, &cfg, db)
}