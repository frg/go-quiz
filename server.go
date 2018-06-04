package main

import (
	"quiz/routes"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	server := echo.New()

	// Gzip
	server.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	server.Use(mw.Logger())
	server.Use(mw.Recover())
	server.Use(mw.CORS())

	// Routes
	server.File("/", "client/dist/index.html")
	server.Static("/", "client/dist/")
	routes.RegisterQuizRoutes(server)

	server.Start(":8000")
}
