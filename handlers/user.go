package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	return c.String(http.StatusOK, "Ini Halaman Register")
}