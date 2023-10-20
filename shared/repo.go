package shared

import "gorm.io/gorm"

type PostgresRepo struct {
	Client *gorm.DB
}
