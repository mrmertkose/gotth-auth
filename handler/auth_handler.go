package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/pkg/password"
	"github.com/mrmertkose/gotth-auth/pkg/sessions"
	"github.com/mrmertkose/gotth-auth/pkg/validation"
	userRepo "github.com/mrmertkose/gotth-auth/repository"
	"github.com/mrmertkose/gotth-auth/request"
	"github.com/mrmertkose/gotth-auth/views/pages"
)

type AuthHandler struct {
	userRepository userRepo.UserRepositoryInterface
}

func AuthHandlerNew() *AuthHandler {
	return &AuthHandler{
		userRepository: userRepo.UserRepositoryNew(),
	}
}

func (a *AuthHandler) Login(c echo.Context) error {
	return render(c, pages.Login())
}

func (a *AuthHandler) PostLogin(c echo.Context) error {
	params := request.LoginRequest{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	responseError := make(map[string]string)
	if err := validator.New().Struct(&params); err != nil {
		responseError = validation.ValidationErrors(err)
		return render(c, pages.LoginForm(params, responseError))
	}

	user, err := a.userRepository.GetUserByEmail(params.Email)
	if err != nil {
		responseError["globalError"] = "E-mail not found."
		return render(c, pages.LoginForm(params, responseError))
	}

	if !password.IsValidPassword(user.Password, params.Password) {
		responseError["globalError"] = "Password incorrect."
		return render(c, pages.LoginForm(params, responseError))
	}

	sessions.SetAuthSession(c, user)

	return hxRedirect(c, "/app/home")
}

func (a *AuthHandler) PostLogout(c echo.Context) error {
	sessions.DeleteAuthSession(c)
	return hxRedirect(c, "/auth/login")
}
