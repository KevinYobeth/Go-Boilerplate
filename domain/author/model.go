package author

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
