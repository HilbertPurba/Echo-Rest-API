package routes

import (
	"net/http"

	"github.com/hilbertpurba/PBP/tugas-crud-echo/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World from Echo!")
	})

	appGroup := e.Group("/users")
	appGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
		TokenLookup:   "cookie:Cookie",
	}))
	appGroup.GET("", controllers.GetAllUsers)
	appGroup.POST("", controllers.InsertNewUser)
	appGroup.PUT("", controllers.UpdateUser)
	appGroup.DELETE("", controllers.DeleteUser)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.Login)
	e.GET("/logout", controllers.Logout)
	return e
}
