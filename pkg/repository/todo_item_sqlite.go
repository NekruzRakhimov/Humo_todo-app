package repository

import (
	"Humo_todo-app/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TodoItemSqlite struct {
	db *sqlx.DB
}

func NewTodoItemSqlite(db *sqlx.DB) *TodoItemSqlite {
	return &TodoItemSqlite{db: db}
}

func (r *TodoItemSqlite) Create(listId int64, item models.TodoItem) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (($1), ($2))", todoItemsTable)
	row, err := tx.Exec(createItemQuery, item.Title, item.Description)
	itemId, err := row.LastInsertId()
	if err != nil {
		err = tx.Rollback()
		return 0, err
	}
	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES (($1), ($2))", listsItemsTable)
	_, err = tx.Exec(createListsItemsQuery, listId, itemId)
	if err != nil {
		err = tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}

func (r *TodoItemSqlite) GetAll(userId, listId int64) ([]models.TodoItem, error) {
	var items []models.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemSqlite) GetById(userId, itemId int64) (models.TodoItem, error) {
	var item models.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemSqlite) Delete(listId, itemId int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	DeletingFromTodoItemsTableQuery := fmt.Sprintf("DELETE FROM %s  WHERE id=$1", todoItemsTable)
	_, err = r.db.Exec(DeletingFromTodoItemsTableQuery, itemId)
	if err != nil {
		err = tx.Rollback()
		return err
	}

	DeletingFromListsItemsTableQuery := fmt.Sprintf("DELETE FROM %s  WHERE list_id=$1 AND item_id=$2", listsItemsTable)
	_, err = r.db.Exec(DeletingFromListsItemsTableQuery, listId, itemId)
	if err != nil {
		err = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *TodoItemSqlite) Update(ListId int64, itemId int64, input models.UpdateItemData) error {
	if input.Title == nil && input.Description == nil {
		query := fmt.Sprintf("UPDATE %s SET done=$1 WHERE id=$2", todoItemsTable)
		_, err := r.db.Exec(query, input.Done, itemId)
		return err
	}

	if input.Title == nil && input.Done == nil {
		query := fmt.Sprintf("UPDATE %s SET description=$1 WHERE id=$2", todoItemsTable)
		_, err := r.db.Exec(query, input.Description, itemId)
		return err
	}

	if input.Description == nil && input.Done == nil {
		query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", todoItemsTable)
		_, err := r.db.Exec(query, input.Title, itemId)
		return err
	}

	if input.Done == nil {
		query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3", todoItemsTable)
		_, err := r.db.Exec(query, input.Title, input.Description, itemId)
		return err
	}

	if input.Title == nil {
		query := fmt.Sprintf("UPDATE %s SET description=$1, done=$2 WHERE id=$3", todoItemsTable)
		_, err := r.db.Exec(query, input.Description, input.Done, itemId)
		return err
	}

	if input.Description == nil {
		query := fmt.Sprintf("UPDATE %s SET title=$1, done=$2 WHERE id=$3", todoItemsTable)
		_, err := r.db.Exec(query, input.Title, input.Done, itemId)
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, done=$3 WHERE id=$4",
		todoItemsTable)
	_, err := r.db.Exec(query, input.Title, input.Description, input.Done, itemId)
	return err
}
