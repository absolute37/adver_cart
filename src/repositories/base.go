package repositories

import "gorm.io/gorm"

type BaseRepo struct {
	GromDB *gorm.DB
}
