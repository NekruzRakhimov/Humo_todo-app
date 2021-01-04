package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-app/models"
)

type TodoListSqlite struct {
	db *sqlx.DB
}

func NewTodoListSqlite(db *sqlx.DB) *TodoListSqlite {
	return &TodoListSqlite{db: db}
}

func (r *TodoListSqlite) Create(userId int64, list models.TodoList) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (($1), ($2))", todoListsTable)
	row, err := tx.Exec(createListQuery, list.Title, list.Description)
	id, err := row.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, nil
	}
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES (($1), ($2))", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}
	return id, tx.Commit()
}

func (r *TodoListSqlite) GetAll(userId int64) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = ($1)",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err

}

func (r *TodoListSqlite) GetById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListSqlite) Delete(userId, listId int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	DeletingFromTodoListsTableQuery := fmt.Sprintf("DELETE FROM %s  WHERE id=$1", todoListsTable)
	_, err = r.db.Exec(DeletingFromTodoListsTableQuery, listId)
	if err != nil {
		err = tx.Rollback()
		return err
	}

	DeletingFromUsersListsTableQuery := fmt.Sprintf("DELETE FROM %s  WHERE user_id=$1 AND list_id=$2", usersListsTable)
	_, err = r.db.Exec(DeletingFromUsersListsTableQuery, userId, listId)
	if err != nil {
		err = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *TodoListSqlite) Update(userId int64, ListId int64, input models.UpdateListData) error{
	if input.Title == nil {
		query := fmt.Sprintf("UPDATE %s SET description=$1 WHERE id=$2", todoListsTable)
		_, err := r.db.Exec(query, input.Description, ListId)
		return err
	}

	if input.Description == nil {
		query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", todoListsTable)
		_, err := r.db.Exec(query, input.Title, ListId)
		return err
	}
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2  WHERE id=$3",
		todoListsTable)
	_, err := r.db.Exec(query, input.Title, input.Description, ListId)
	return err
}