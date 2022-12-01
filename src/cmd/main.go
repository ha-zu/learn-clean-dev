package main

import (
	"net/http"

	cnf "github.com/ha-zu/learn-clean-dev/src/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			cnf.Env.DevURL,
		},
		AllowMethods: []string{
			http.MethodGet,
		},
	}))

	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusOK,
			Message: "This is a health check.",
		})
	})

	//ToDo setting port statement
	e.Logger.Fatal(e.Start(":" + cnf.Env.DevPort))
}
