package models

type LineItems struct {
	UUID     string  `json:"uuid"`
	Name     string  `json:"name" validate:"min=5"`
	Price    float64 `json:"price" validate:"gt=0"`
	MenuUuid string  `json:"menuUuid" validate:"uuid4"`
}
