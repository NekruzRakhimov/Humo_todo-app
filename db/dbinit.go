package db

import (
	"github.com/jmoiron/sqlx"
	"log"
)

// Init Initializing tables
func Init(database *sqlx.DB) {
	DDLs := []string{
		CreateUsersTable,
		CreateTodoListsTable,
		CreateUsersListsTable,
		CreateTodoItemsTable,
		CreateListsItemsTable,
	}

	for _, ddl := range DDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatal("Error while creating table. Error is:", err)
		}
	}
}

//Dropping tables
func Drop(database *sqlx.DB) {
	DDLs := []string{
		DropListsItemsTable,
		DropUsersListsTable,
		DropTodoListsTable,
		DropUsersTable,
		DropTodoItemsTable,
	}

	for _, ddl := range DDLs {
		_, err := database.Exec(ddl)
		if err != nil {
			log.Fatal("Error while dropping table. Error is:", err)
		}
	}
}
