package handler

import (
	"net/http"

	"github.com/fignocius/rfid-api/view/model"
	"github.com/fignocius/rfid-api/view/service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

// Product struct of handler
type Product struct {
	service.ProductService
}

// NewProductHandler is a function to instance a new ProductHandler
func NewProductHandler(db *sqlx.DB) ProductHandler {
	return &Product{
		service.New(db),
	}
}

// View is a function to View the Product status
func (t Product) View(c echo.Context) error {
	request := model.ProductRequest{}
	err := c.Bind(request)
	if err != nil {
		return echo.ErrBadRequest
	}
	products, err := t.ProductService.Find(request)
	if err != nil {
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, products)
}
