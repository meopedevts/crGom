package handler

import (
	"github.com/labstack/echo/v4"
	// "github.com/meopedevts/crgom/model"
	"github.com/meopedevts/crgom/view/register"
)

type RegisterHandler struct{}

func (h RegisterHandler) HandleRegisterPage(c echo.Context) error {
	return render(c, register.RegisterForm(register.RegisterFormProps{
		Context: c,
	}))
}

// func (h RegisterHandler) HandleRegisterPost(c echo.Context) error {
// 	model.NewUser(
// 		c.FormValue("firstName"),
// 		c.FormValue("lastName"),
// 		c.FormValue("email"),
// 	)
// 	return render(c, register.RegisterSuccess())
// }
