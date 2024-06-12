package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func hxRedirect(c echo.Context, to string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", to)
		return c.NoContent(http.StatusSeeOther)
	}
	return c.Redirect(http.StatusSeeOther, to)
}
