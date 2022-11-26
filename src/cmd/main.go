package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load("src/configs/.env")
	if err != nil {
		log.Fatalf("Error Configer Loading env %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			os.Getenv("LOCALHOST"),
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

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
