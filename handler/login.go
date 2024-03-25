package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/meopedevts/crgom/model"
	"github.com/meopedevts/crgom/view/login"
)

type LoginHandler struct{}

func (h LoginHandler) HandleLoginPage(c echo.Context) error {
	user := model.GetUser()
	return render(c, login.LoginForm(user))
}
