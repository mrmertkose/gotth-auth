package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/pkg/sessions"
	"net/http"
)

func IsGuest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuth := sessions.GetIsAuth(c)
		if isAuth == true {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}
