package repository

import (
	"github.com/jackc/pgx/pgtype"
	"github.com/jmoiron/sqlx"
	"test/internal/features/products/models"
)

type ProductsRepository struct {
	Db *sqlx.DB
}

func NewProductsRepository(db *sqlx.DB) *ProductsRepository {
	return &ProductsRepository{Db: db}
}

func (ProductsRepository) GetByID(id pgtype.UUID) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (ProductsRepository) GetList(limit, offset int) (*models.GetProductsResponse, error) {
	//TODO implement me
	panic("implement me")
}
