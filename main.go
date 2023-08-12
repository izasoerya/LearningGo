package main

import (
	"github.com/izasoerya/Prakerja-Go/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	e.GET("/siji", controller.Siji)
	e.GET("/loro", controller.Loro)
	e.GET("/telu", controller.Telu)
	e.GET("/papat", controller.Papat)

	e.Logger.Fatal(e.Start(":2000"))
}


