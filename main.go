package main

import (
	"context"
	"log"
	"os"

	"github.com/Alohadancemeow/go-sekai-shop-demo/config"
	database "github.com/Alohadancemeow/go-sekai-shop-demo/pkg/database/script"
)

func main() {
	ctx := context.Background()

	// Initialize config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())

	// Database connection
	db := database.DbConn(ctx, &cfg)
	log.Println(db)
	defer db.Disconnect(ctx)


	// Start Server
	// server.Start(ctx, &cfg, db)
}