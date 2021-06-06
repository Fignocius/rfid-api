package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/rfid-api/view/model"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// ProductRepository is a interface of repository
type ProductRepository interface {
	Get(code string) (model.Product, error)
}
