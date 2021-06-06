package model

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Code      string     `json:"code" db:"code"`
	Name      string     `json:"name" db:"name"`
	Value     int64      `json:"value" db:"value"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

func (m Product) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return string("")
	}
	return string(b)
}

type ProductRequest struct {
	Codes []string `query:"code"`
}

type ProductResponse struct {
	Product    Product `json:"product"`
	Quantity   int64   `json:"quantity"`
	TotalValue int64   `json:"total_value"`
}

func (m ProductResponse) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return string("")
	}
	return string(b)
}

func (m *ProductResponse) CalcValue() {
	m.TotalValue = m.Quantity * m.Product.Value
}

func (m *ProductResponse) IncrementQtd() {
	m.Quantity = m.Quantity + 1
}
