package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/views/pages"
)

type AppHandler struct {
}

func AppHandlerNew() *AppHandler {
	return &AppHandler{}
}

func (a *AppHandler) Index(c echo.Context) error {
	return render(c, pages.Index())
}

func (a *AppHandler) App(c echo.Context) error {
	return render(c, pages.App())
}
