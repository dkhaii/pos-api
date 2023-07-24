package main

import (
	"strconv"

	"github.com/gerinmordekhai/pos-api/handlers"
	"github.com/labstack/echo/v4"
)

const PORT = 8080

func main() {
	router := echo.New()

	router.GET("/", handlers.RegisterUser)

	router.Logger.Fatal(router.Start(":" + strconv.Itoa(PORT)))
}