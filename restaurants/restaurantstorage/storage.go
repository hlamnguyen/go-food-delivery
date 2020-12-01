package restaurantstorage

import "gorm.io/gorm"

type storeMySQL struct {
	db *gorm.DB
}

func NewMySQLStore(db *gorm.DB) *storeMySQL {
	return &storeMySQL{db: db}
}
