package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/pkg/sessions"
)

func AuthContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAuth := sessions.GetIsAuth(c)
		id := sessions.GetId(c)
		name := sessions.GetName(c)
		c.Set("isAuth", isAuth)
		c.Set("id", id)
		c.Set("name", name)
		return next(c)
	}
}
