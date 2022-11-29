package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ha-zu/learn-clean-dev/src/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	configs.EnvSetting()
}

func main() {

	cnf, err := configs.New()
	if err != nil {
		log.Fatalf("Configer New Error: %v", err)
		os.Exit(1)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			cnf.DevUrl,
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
	e.Logger.Fatal(e.Start(":" + cnf.DevPort))
}
