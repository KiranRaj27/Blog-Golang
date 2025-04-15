// cmd/blog-api/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiranraj27/blog-golang/internal/config"
	"github.com/kiranraj27/blog-golang/internal/router"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Initialize DB
	config.InitDB(cfg)

	// Initialize router
	r := router.New()

	// Start HTTP server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server is running on %s\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
