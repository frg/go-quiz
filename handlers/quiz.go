package handlers

import (
	"net/http"

	"quiz/models"

	"github.com/labstack/echo"
)

func GetQuiz() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetQuiz())
	}
}

func SaveUserAnswer(user string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": 999,
		})
	}
}

func GetQuizLeaderboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "tasks")
	}
}

func GetRelativeQuizLeaderboard(user string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "tasks")
	}
}
