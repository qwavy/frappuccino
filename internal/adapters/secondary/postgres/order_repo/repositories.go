package order_repo

import (
	"context"
	"frappucchino/internal/core/domain/model/order"
	"frappucchino/internal/core/ports"

	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) ports.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll(ctx context.Context) ([]*order.Order, error) {
	rows, err := r.db.Query(ctx, `SELECT id, customerName, items, status, createdAt FROM order`)
	if err != nil {
		return nil, err
	}

	var orders []*order.Order

	for rows.Next() {
		var orderDTO OrderDTO

		err = rows.Scan(&orderDTO.Id, &orderDTO.CustomerName, &orderDTO.Items, &orderDTO.Status, &orderDTO.CreatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, DTOtoDomain(orderDTO))
	}

	if rows.Err() != nil {
		return nil, err
	}

	return orders, err
}

func (r *orderRepository) GetById(ctx context.Context, id string) (*order.Order, error) {
	row := r.db.QueryRow(ctx, `SELECT id, customerName, items, status, createdAt FROM order WHERE id = $1`, string)

	var orderDTO OrderDTO
	err := row.Scan(&orderDTO.Id, &orderDTO.CustomerName, &orderDTO.Items, &orderDTO.Status, &orderDTO.CreatedAt)
	if err != nil {
		return nil, err
	}

	return DTOtoDomain(orderDTO), nil
}

func (r *orderRepository) Create(ctx context.Context, order *order.Order) error {
	orderDTO := DomainToDTO(order)
	_, err := r.db.Exec(ctx, `INSERT INTO menu (id, customerName, items, status, createdAt) values ($1, $2, $3, $4, $5)`, orderDTO.Id, orderDTO.CustomerName, orderDTO.Items, orderDTO.Status, orderDTO.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) DeleteById(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM order WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) UpdateById(ctx context.Context, id string, order *order.Order) error {
	orderDTO := DomainToDTO(order)
	_, err := r.db.Exec(ctx, `UPDATE order SET customerName = $2, items = $3, status = $4, createdAt = $5 WHERE id = $1`, id, orderDTO.CustomerName, orderDTO.Items, orderDTO.Status, orderDTO.CreatedAt)
	if err != nil {
		return err
	}

	return err
}
