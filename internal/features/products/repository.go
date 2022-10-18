package products

import (
	"github.com/jackc/pgx/pgtype"
	"test/internal/features/products/models"
)

type Repository interface {
	GetByID(id pgtype.UUID) (*models.Product, error)
	GetList(limit, offset int) (*models.GetProductsResponse, error)
}
