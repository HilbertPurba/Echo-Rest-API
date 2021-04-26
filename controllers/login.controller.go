package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hilbertpurba/PBP/tugas-crud-echo/helpers"
	"github.com/hilbertpurba/PBP/tugas-crud-echo/models"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	Level string `json:"level"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := models.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	expiryTime := time.Now().Add(2 * time.Hour)
	token, err := generateToken(email, expiryTime)
	cookie := &http.Cookie{
		Name:    "Cookie",
		Value:   token,
		Expires: expiryTime,
	}

	c.SetCookie(cookie)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login Success! Welcome " + email + "!",
		"token":   token,
	})

}

func generateToken(email string, expiryTime time.Time) (string, error) {

	claims := jwtCustomClaims{
		Email: email,
		Level: "application",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
