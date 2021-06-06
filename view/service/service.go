package service

import (
	"github.com/fignocius/rfid-api/view/model"
)

// ProductService is a interface of service
type ProductService interface {
	Find(products model.ProductRequest) ([]model.ProductResponse, error)
}
