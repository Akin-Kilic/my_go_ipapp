package main

import (
	"log"
	"my_go_project/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/ip", handlers.GetIPInfo)

	// Start server
	log.Fatal(e.Start(":8080"))
}
