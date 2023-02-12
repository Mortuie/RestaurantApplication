package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDbClient(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	return db, err
}
