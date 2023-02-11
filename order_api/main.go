package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	UUID  string `gorm:"primaryKey"`
	Code  string
	Price uint
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.UUID = uuid.NewString()
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to sqlitedb")
	}

	db.AutoMigrate(&Product{})
	// db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product)
	// db.Delete(&product)

	fmt.Println(product)
}
