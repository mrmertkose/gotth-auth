package sessions

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/mrmertkose/gotth-auth/model"
	"net/http"
)

var SessionName = "session"
var secretKey = "asd6935+"

func Start(app *echo.Echo) {
	app.Use(session.Middleware(sessions.NewCookieStore([]byte(secretKey))))
}

func GetIsAuth(c echo.Context) bool {
	sess, _ := session.Get(SessionName, c)
	response, ok := sess.Values["isAuth"]
	if !ok {
		return false
	}
	return response.(bool)
}

func GetId(c echo.Context) uint {
	sess, _ := session.Get(SessionName, c)
	response, ok := sess.Values["id"]
	if !ok {
		return 0
	}
	return response.(uint)
}

func GetName(c echo.Context) string {
	sess, _ := session.Get(SessionName, c)
	response, ok := sess.Values["name"]
	if !ok {
		return ""
	}
	return response.(string)
}

func SetAuthSession(c echo.Context, user *model.User) {
	sess, _ := session.Get(SessionName, c)
	sess.Options.SameSite = http.SameSiteLaxMode
	sess.Options.Secure = true
	sess.Options.HttpOnly = true
	sess.Options.MaxAge = 3600
	sess.Values["isAuth"] = true
	sess.Values["id"] = user.Id
	sess.Values["name"] = user.Name
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return
	}
}

func DeleteAuthSession(c echo.Context) {
	sess, _ := session.Get(SessionName, c)
	sess.Options.MaxAge = -1
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		return
	}
}
