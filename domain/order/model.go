package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId     uint64     `json:"order_id" gorm:"primaryKey"`
	CustomerId  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items" gorm:"foreignKey:OrderID"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemId   uuid.UUID `json:"item_id" gorm:"primaryKey"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
