package models

type Menus struct {
	UUID           string `json:"uuid"`
	Name           string `json:"name" validate:"min=5"`
	RestaurantUuid string `json:"restaurantUuid" validate:"min=5"`
}
