package models

import "golang.org/x/crypto/bcrypt"

type UserResponse struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username" validate:"min=3"`
	Password string `json:"password" validate:"min=3"`
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
