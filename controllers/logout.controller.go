package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("Cookie")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie.Value = ""
	cookie.Expires = time.Now()
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logout Success! Goodbye!",
	})
}
