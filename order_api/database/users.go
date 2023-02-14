package database

import (
	"fmt"
	"restapp/order-api/models"
)

func InsertUser(d *DB, u models.User) error {

	statement := fmt.Sprintf(`insert into users (uuid, username, password) values ('%s', '%s', '%s')`, u.UUID, u.Username, u.Password)

	_, err := d.db.Exec(statement)
	return err
}

func GetUser(d *DB, uuid string) (*models.UserResponse, error) {
	statement := fmt.Sprintf(`select uuid, username from users where uuid = '%s';`, uuid)

	row := d.db.QueryRow(statement)

	var u models.UserResponse
	err := row.Scan(&u.UUID, &u.Username)
	if err != nil {
		return &models.UserResponse{}, err
	}

	return &u, nil
}
