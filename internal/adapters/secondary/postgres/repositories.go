package postgres

import (
	"frappucchino/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	db                  *pgxpool.Pool
	InventoryRepository ports.InventoryRepository
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{db: db}
}
