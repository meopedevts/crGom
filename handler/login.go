package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/meopedevts/crgom/view/login"
)

type LoginHandler struct{}

func (h LoginHandler) HandleLoginPage(c echo.Context) error {
	return render(c, login.LoginForm(login.LoginFormProps{
		Context: c,
	}))
}
