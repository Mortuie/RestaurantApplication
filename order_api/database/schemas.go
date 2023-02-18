package database

func (db *DB) getSchemas() map[string]string {
	return map[string]string{
		"users": `create table if not exists users (
			uuid string not null primary key,
			username string not null unique,
			password string not null
		);`,
		"restaurants": `create table if not exists restaurants (
			uuid string not null primary key,
			name string not null,
			user_uuid string not null,
			foreign key(user_uuid) references users(uuid)
		)`,
		"menus": `create table if not exists menus (
			uuid string not null primary key,
			name string not null,
			restaurant_uuid string not null,
			foreign key(restaurant_uuid) references restaurants(uuid)
		)`,
		"menu_categories": `create table if not exists menu_categories (
			uuid string not null primary key,
			category_name string not null,
			menu_uuid string not null,
			foreign key(menu_uuid) references menus(uuid)
		)`,
		"line_items": `create table if not exists line_items (
			uuid string not null primary key,
			name string not null,
			price float not null,
			menu_uuid string not null,
			foreign key(menu_uuid) references menus(uuid)
		)`,
	}
}
