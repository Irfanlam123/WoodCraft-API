package main

import (
	"fmt"
	"log"

	"WoodCraft-API/config"
	"WoodCraft-API/database"
	"WoodCraft-API/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load Config (env variables)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to Database
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Create Echo instance
	e := echo.New()

	// ✅ Add CORS Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"}, // Frontend origin
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Register Routes
	routes.RegisterRoutes(e, db)

	// Start Server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server running on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}
}
