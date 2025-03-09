package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	e.GET("/200", func(c echo.Context) error {
		// return c.HTML(http.StatusOK, "Hello, Docker! <3")
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	e.GET("/400", func(c echo.Context) error {
		// return c.HTML(http.StatusOK, "Hello, Docker! <3")
		return c.JSON(http.StatusBadRequest, struct{ Status string }{Status: "BadRequest"})
	})
	e.GET("/500", func(c echo.Context) error {
		// return c.HTML(http.StatusOK, "Hello, Docker! <3")
		return c.JSON(http.StatusInternalServerError, struct{ Status string }{Status: "InternalServerError"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
