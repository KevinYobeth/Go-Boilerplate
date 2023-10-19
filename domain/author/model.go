package author

import "github.com/google/uuid"

type Author struct {
	Id   uuid.UUID `json:"id" gorm:"primaryKey"`
	Name string    `json:"name"`
}
