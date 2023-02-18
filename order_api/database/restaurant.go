package database

import (
	"fmt"
	"restapp/order-api/models"
)

func InsertRestaurant(d *DB, r models.Restaurant) error {
	_, err := d.db.Exec("insert into restaurants (uuid, name, user_uuid) values (?, ?, ?);", r.UUID, r.Name, r.UserUuid)
	return err
}

func GetRestaurant(d *DB, uuid string) (*models.Restaurant, error) {
	row := d.db.QueryRow("select uuid, name, user_uuid from restaurants where uuid = ?;", uuid)

	var r models.Restaurant
	if err := row.Scan(&r.UUID, &r.Name, &r.UserUuid); err != nil {
		return &models.Restaurant{}, err
	}

	return &r, nil
}

func GetRestaurants(d *DB, pc models.PagingConfig) ([]models.Restaurant, error) {
	rows, err := d.db.Query("select uuid, name, user_uuid from restaurants limit ? offset ?;", pc.PageSize, pc.Offset)

	if err != nil {
		return []models.Restaurant{}, nil
	}

	var rs []models.Restaurant
	for rows.Next() {
		var r models.Restaurant
		if err := rows.Scan(&r.UUID, &r.Name, &r.UserUuid); err != nil {
			return rs, err
		}
		rs = append(rs, r)
	}
	fmt.Println("ROWS", rs)

	if err = rows.Err(); err != nil {
		return rs, err
	}

	return rs, nil
}
