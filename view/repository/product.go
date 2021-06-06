package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fignocius/rfid-api/view/model"
	"github.com/jmoiron/sqlx"
)

// Product struct repository
type Product struct {
	db *sqlx.DB
}

// NewProduct function to instance new Product
func NewProduct(db *sqlx.DB) ProductRepository {
	return &Product{
		db: db,
	}
}

// Get function to get fields of the Product
func (t Product) Get(code string) (product model.Product, err error) {
	query := psql.Select("*").
		From("product").
		Where(sq.Eq{"code": code})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := t.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Get(&product, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}
