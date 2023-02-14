package database

func (db *DB) getSchemas() map[string]string {
	return map[string]string{
		"users": `create table if not exists users (uuid string not null primary key, username string not null unique, password string not null);`,
		// "restaurants":     ``,
		// "menus":           ``,
		// "menu_line_items": ``,
		// "menu_categories": ``,
	}
}
