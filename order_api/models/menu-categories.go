package models

type MenuCategory struct {
	UUID         string `json:"uuid"`
	CategoryName string `json:"categoryName" validate:"min=5"`
	MenuUuid     string `json:"menuUuid" validate:"uuid4"`
}
