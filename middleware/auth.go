package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/pkg/sessions"
	"github.com/mrmertkose/gotth-auth/repository"
	"net/http"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	var userRepo = repository.UserRepositoryNew()
	return func(c echo.Context) error {
		isAuth := sessions.GetIsAuth(c)
		if isAuth == false {
			return c.Redirect(http.StatusSeeOther, "/auth/login")
		}

		user, err := userRepo.GetUserByID(sessions.GetId(c))
		if err != nil || user.Id == 0 {
			return c.Redirect(http.StatusSeeOther, "/auth/login")
		}

		return next(c)
	}
}
