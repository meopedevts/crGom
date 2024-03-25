package main

import (
	"github.com/labstack/echo/v4"
	"github.com/meopedevts/crgom/handler"
)

func main() {
	e := echo.New()

	registerHandler := handler.RegisterHandler{}
	loginHandler := handler.LoginHandler{}
	e.Static("/static", "static")
	e.GET("/register", registerHandler.HandleRegisterPage)
	e.POST("/register", registerHandler.HandleRegisterPost)
	e.GET("/login", loginHandler.HandleLoginPage)

	e.Start(":8080")
}
