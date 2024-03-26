package main

import (
	"github.com/labstack/echo/v4"
	"github.com/meopedevts/crgom/handler"
)

func main() {
	e := echo.New()

	// Handlers
	registerHandler := handler.RegisterHandler{}
	loginHandler := handler.LoginHandler{}
	homeHandler := handler.HomeHandler{}
	//

	// Routes
	e.Static("/static", "static")
	e.GET("/auth/register", registerHandler.HandleRegisterPage)
	// e.POST("/register", registerHandler.HandleRegisterPost)
	e.GET("/auth/login", loginHandler.HandleLoginPage)
	e.GET("/", homeHandler.HandleHomePage)
	//

	e.Start(":8080")
}
