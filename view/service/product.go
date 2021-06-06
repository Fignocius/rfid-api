package service

import (
	"strings"

	"github.com/fignocius/rfid-api/view/model"
	"github.com/fignocius/rfid-api/view/repository"
	"github.com/jmoiron/sqlx"
)

// Product struct repository
type Product struct {
	Repository repository.ProductRepository
}

// NewProduct function to instance new Product
func New(db *sqlx.DB) ProductService {
	return &Product{
		Repository: repository.NewProduct(db),
	}
}

// Find function to find informations of the Products
func (t Product) Find(request model.ProductRequest) (products []model.ProductResponse, err error) {
	var prod model.Product

	for i, code := range request.Codes {
		var hasCode bool

		if i == 0 {
			prod, err = t.Repository.Get(code)
			if err != nil {
				return
			}
			newProduct := model.ProductResponse{Product: prod, Quantity: 1}
			newProduct.CalcValue()
			products = append(products, newProduct)
			continue
		}
		for _, c := range products {
			if strings.EqualFold(code, c.Product.Code) {
				c.Quantity++
				c.CalcValue()
				hasCode = true
			}
		}
		if !hasCode {
			prod, err = t.Repository.Get(code)
			if err != nil {
				return
			}
			newProduct := model.ProductResponse{Product: prod, Quantity: 1}
			newProduct.CalcValue()
			products = append(products, newProduct)
		}
	}
	return
}
