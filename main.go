package main

import (
	"net/http"

	// soalprakerja "github.com/izasoerya/Prakerja-Go/soal_prakerja"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":2000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}
