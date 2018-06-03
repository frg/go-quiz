package main

import (
	"quiz/routes"

	"github.com/hb-go/echo-web/module/cache"
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

	// Cache
	server.Use(cache.Cache())

	// Routes
	server.File("/", "client/index.html")
	routes.RegisterQuizRoutes(server)

	server.Start(":8000")
}
