package db

//Creating tables
const (
	CreateUsersTable = `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
	);`

	CreateTodoListsTable = `CREATE TABLE IF NOT EXISTS todo_lists(
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    title TEXT NOT NULL,
    description TEXT NOT NULL
	);`

	CreateUsersListsTable = `CREATE TABLE IF NOT EXISTS users_lists(
    id INTEGER PRIMARY KEY  AUTOINCREMENT UNIQUE,
    user_id INTEGER REFERENCES users(id) NOT NULL,
    list_id INTEGER REFERENCES todo_lists(id) NOT NULL
	);`

	CreateTodoItemsTable = `CREATE TABLE IF NOT EXISTS todo_items(
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
	);`

	CreateListsItemsTable = `CREATE TABLE IF NOT EXISTS lists_items(
    id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    item_id INTEGER REFERENCES todo_items(id) NOT NULL,
    list_id INTEGER REFERENCES todo_lists(id) NOT NULL
	);`
)

//Dropping tables
const (
	DropListsItemsTable = `DROP TABLE IF EXISTS lists_items`

	DropUsersListsTable = `DROP TABLE IF EXISTS users_lists`

	DropTodoListsTable  = `DROP TABLE IF EXISTS todo_lists`

	DropUsersTable      = `DROP TABLE IF EXISTS users`

	DropTodoItemsTable  = `DROP TABLE IF EXISTS todo_items`
)

