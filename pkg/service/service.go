package service

import (
	"todo-app/models"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type TodoList interface {
	Create(userId int64, list models.TodoList) (int64, error)
	GetAll(userId int64) ([]models.TodoList, error)
	GetById(userId, listId int64) (models.TodoList, error)
	Update(userId int64, ListId int64, input models.UpdateListData) error
	Delete(userId, listId int64) error
}

type TodoItem interface {
	Create(userId, listId int64, item models.TodoItem) (int64, error)
	GetAll(userId, listId int64) ([]models.TodoItem, error)
	GetById(userId, itemId int64) (models.TodoItem, error)
	Update(ListId int64, itemId int64, input models.UpdateItemData) error
	Delete(listId, itemId int64) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}