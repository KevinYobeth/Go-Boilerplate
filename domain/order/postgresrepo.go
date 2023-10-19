package order

import (
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type PostgresRepo struct {
	Client *gorm.DB
}

func (r *PostgresRepo) Insert(ctx context.Context, order Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	fmt.Println(data)

	result := r.Client.Create(&data)

	if result.Error != nil {
		return fmt.Errorf("failed to add to orders: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) FindById(ctx context.Context, orderId uint64) (Order, error) {
	return Order{}, nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, orderId uint64) error {
	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, order Order) error {
	return nil
}

func (r *PostgresRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	var orders []Order

	if result := r.Client.Find(&orders); result.Error != nil {
		// fmt.Println(result.Error)
		fmt.Println("ERROR DISINI")
	}

	return FindResult{
		Orders: orders,
		Cursor: 0,
	}, nil
}
