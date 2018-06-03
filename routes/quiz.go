package routes

import (
	"github.com/labstack/echo"
)

// RegisterQuizRoutes Register quiz routes
func RegisterQuizRoutes(server *echo.Echo) *echo.Echo {
	// GET quiz questions & answers
	server.GET("/quiz", func(context echo.Context) error {
		return context.JSON(200, "GET Quiz")
	})

	// POST quiz answers for user
	server.POST("/quiz/answer/:user", func(context echo.Context) error {
		return context.JSON(200, "POST Quiz '"+context.Param("user")+"'")
	})

	// GET quiz leaderboard
	server.GET("/quiz/leaderboard", func(context echo.Context) error {
		return context.JSON(200, "GET /quiz/leaderboard")
	})

	// GET quiz leaderboard relative to user stats
	server.GET("/quiz/leaderboard/relative/:user", func(context echo.Context) error {
		return context.JSON(200, "GET /quiz/highscore/relative/:user '"+context.Param("user")+"'")
	})

	return server
}
