package router

import (
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/handler"
	"github.com/mrmertkose/gotth-auth/middleware"
)

func Setup(app *echo.Echo) {
	authHandler := handler.AuthHandlerNew()
	appHandler := handler.AppHandlerNew()

	app.GET("/", appHandler.Index)

	// AUTH
	auth := app.Group("/app", middleware.Auth)
	auth.GET("/home", appHandler.App)
	auth.POST("/logout", authHandler.PostLogout)

	// GUEST
	guest := app.Group("/auth", middleware.IsGuest)
	guest.GET("/login", authHandler.Login)
	guest.POST("/login", authHandler.PostLogin)
}
