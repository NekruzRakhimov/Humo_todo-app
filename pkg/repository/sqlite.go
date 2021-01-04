package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	usersTable 	    = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

//Инициализация базы данных
func NewSqliteDB(dbName string) (*sqlx.DB, error) {
	database, err := sqlx.Open("sqlite3", fmt.Sprintf("%s", dbName))
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	return database, nil
}