package main

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	UUID string `gorm:"primaryKey"`
	Name string
}

type Menu struct {
	gorm.Model
	UUID      string `gorm:"primaryKey"`
	Name      string
	LineItems []LineItem
	MenuType  []string
}

type LineItem struct {
	gorm.Model
	UUID  string `gorm:"primaryKey"`
	Name  string
	Type  string
	Price uint
}
