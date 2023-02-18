package models

type Restaurant struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name" validate:"min=5"`
	UserUuid string `json:"userUuid" validate:"min=5"`
}
