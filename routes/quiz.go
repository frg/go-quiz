package routes

import (
	"quiz/handlers"

	"github.com/labstack/echo"
)

// RegisterQuizRoutes Register quiz routes
func RegisterQuizRoutes(server *echo.Echo) *echo.Echo {
	// GET quiz questions & answers
	server.GET("/quiz", handlers.GetQuiz())

	// POST quiz answers for user
	server.POST("/quiz/answer", handlers.SaveUserAnswer())

	// GET quiz leaderboard
	server.GET("/quiz/leaderboard", handlers.GetQuizLeaderboard())

	// GET quiz leaderboard relative to user stats
	server.GET("/quiz/leaderboard/relative/:user", handlers.GetRelativeQuizLeaderboard())

	return server
}
