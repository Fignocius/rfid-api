package handler

import (
	"github.com/labstack/echo/v4"
)

// ProductHandler interface
type ProductHandler interface {
	View(c echo.Context) error
}
