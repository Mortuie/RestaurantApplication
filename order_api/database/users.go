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

func GetUsers(d *DB, pc models.PagingConfig) ([]models.UserResponse, error) {
	statement := fmt.Sprintf(`select uuid, username from users limit %d offset %d;`, pc.PageSize, pc.Offset)

	rows, err := d.db.Query(statement)

	if err != nil {
		return []models.UserResponse{}, err
	}

	var users []models.UserResponse
	for rows.Next() {
		var u models.UserResponse
		if err := rows.Scan(&u.UUID, &u.Username); err != nil {
			return users, err
		}
		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}
