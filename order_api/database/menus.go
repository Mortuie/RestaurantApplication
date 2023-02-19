package database

import "restapp/order-api/models"

func InsertMenu(d *DB, m models.Menus) error {
	_, err := d.db.Exec("insert into restaurants (uuid, name, restaurant_uuid) values (?, ?, ?);", m.UUID, m.Name, m.RestaurantUuid)
	return err
}

func GetMenu(d *DB, uuid string) (*models.Menus, error) {
	row := d.db.QueryRow("select uuid, name, restaurant_uuid from menus where uuid = ?;", uuid)

	var m models.Menus
	if err := row.Scan(&m.UUID, &m.Name, &m.RestaurantUuid); err != nil {
		return &models.Menus{}, err
	}

	return &m, nil
}

func GetMenus(d *DB, pc models.PagingConfig) ([]models.Menus, error) {
	rows, err := d.db.Query("select uuid, name, restaurant_uuid from menus limit ? offset ?;", pc.PageSize, pc.Offset)

	if err != nil {
		return []models.Menus{}, nil
	}

	var ms []models.Menus
	for rows.Next() {
		var m models.Menus
		if err := rows.Scan(&m.UUID, &m.Name, &m.RestaurantUuid); err != nil {
			return ms, err
		}
		ms = append(ms, m)
	}

	if err = rows.Err(); err != nil {
		return ms, err
	}

	return ms, nil
}
