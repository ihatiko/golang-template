package models

import "github.com/jackc/pgx/pgtype"

type Product struct {
	Id    pgtype.UUID `json:"id" db:"id"`
	Name  string      `json:"name" db:"name"`
	Image string      `json:"image" db:"image"`
	Price float32     `json:"price" db:"price"`
}

type GetProductsResponse struct {
	Data  []*Product `json:"data"`
	Count int        `json:"count"`
}
