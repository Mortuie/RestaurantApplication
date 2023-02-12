package main

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RestaurantUser struct {
	gorm.Model
	Username    string       `gorm:"uniqueIndex" json:"username"`
	Password    string       `json:"password"`
	Restaurants []Restaurant `gorm:"foreignKey:ID" json:"restaurants"`
}

func (ru *RestaurantUser) Validate() error {
	if ru.Username == "" || ru.Password == "" {
		return errors.New("username || password is empty")
	}

	return nil
}

func (ru *RestaurantUser) BeforeCreate(tx *gorm.DB) error {
	b, err := bcrypt.GenerateFromPassword([]byte(ru.Password), 14)

	if err != nil {
		fmt.Println("error hashing password")
		return err
	}
	ru.Password = string(b)

	return nil
}

type Restaurant struct {
	gorm.Model
	Name  string
	Menus []Menu `gorm:"foreignKey:ID"`
}

type Menu struct {
	gorm.Model
	Name            string         `json:"name"`
	OwnedRestaurant uint           `gorm:"foreignKey:ID" json:"ownedRestaurant"`
	LineItems       []LineItem     `gorm:"foreignKey:ID" json:"lineItems"`
	MenuSections    []MenuSections `gorm:"foreignKey:ID" json:"menuSections"`
}

type MenuSections struct {
	gorm.Model
	Name string `json:"name"`
}

func (m *Menu) Validate() error {
	fmt.Println("menu:", m)
	if m.Name == "" {
		fmt.Println("name:", m.Name, "lineItems:", m.LineItems, "menuItems:", m.MenuSections)
		return errors.New("some of the required fields are empty")
	}

	if m.LineItems == nil {
		m.LineItems = []LineItem{}
	}

	if m.MenuSections == nil {
		m.MenuSections = []MenuSections{}
	}

	return nil
}

type LineItem struct {
	gorm.Model
	Name  string
	Type  string
	Price float32
}
