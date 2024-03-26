package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/meopedevts/crgom/view/home"
)

type HomeHandler struct{}

func (a HomeHandler) HandleHomePage(c echo.Context) error {
	return render(c, home.Home(home.HomeProps{
		Context: c,
	}))
}
