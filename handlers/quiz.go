package handlers

import (
	"net/http"

	"quiz/models"

	"github.com/labstack/echo"
)

// GetQuiz Get quiz data
func GetQuiz() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, models.GetQuiz())
	}
}

// SaveUserAnswer Save user answers
func SaveUserAnswer() echo.HandlerFunc {
	return func(context echo.Context) error {
		var sub models.QuizUserSubmission

		// Bind context to sub pointer
		context.Bind(&sub)

		models.SaveUserAnswer(sub)

		return context.JSON(http.StatusCreated, nil)
	}
}

// GetQuizLeaderboard Get quiz leader board
func GetQuizLeaderboard() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, "tasks")
	}
}

// GetRelativeQuizLeaderboard Get quiz user relative stat
func GetRelativeQuizLeaderboard(user string) echo.HandlerFunc {
	return func(context echo.Context) error {
		var user = context.Param("user")

		return context.JSON(http.StatusOK, models.GetRelativeQuizLeaderboard(user))
	}
}
