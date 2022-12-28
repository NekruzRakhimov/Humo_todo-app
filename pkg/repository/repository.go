package repository

import (
	"Humo_todo-app/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int64, list models.TodoList) (int64, error)
	GetAll(userId int64) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int64) error
	Update(userId int64, ListId int64, input models.UpdateListData) error
}

type TodoItem interface {
	Create(listId int64, item models.TodoItem) (int64, error)
	GetAll(userId, listId int64) ([]models.TodoItem, error)
	GetById(userId, itemId int64) (models.TodoItem, error)
	Update(ListId int64, itemId int64, input models.UpdateItemData) error
	Delete(listId, itemId int64) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
		TodoList:      NewTodoListSqlite(db),
		TodoItem:      NewTodoItemSqlite(db),
	}
}
