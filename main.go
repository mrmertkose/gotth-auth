package main

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/middleware"
	"github.com/mrmertkose/gotth-auth/pkg/database"
	customMid "github.com/mrmertkose/gotth-auth/pkg/middleware"
	"github.com/mrmertkose/gotth-auth/pkg/sessions"
	"github.com/mrmertkose/gotth-auth/router"
)

//go:embed all:assets
var StaticFS embed.FS

func main() {
	app := echo.New()
	database.Connect()
	sessions.Start(app)
	app.Static("/", "assets")

	// ** //
	app.Use(customMid.ContextValue)
	app.Use(middleware.AuthContext)

	router.Setup(app)
	app.Logger.Fatal(app.Start(":3000"))
}
