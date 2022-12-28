package service

import (
	"Humo_todo-app/models"
	"Humo_todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int64, list models.TodoList) (int64, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int64) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int64) (models.TodoList, error) {
	return s.repo.GetById(int(userId), int(listId))
}

func (s *TodoListService) Delete(userId, listId int64) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId int64, ListId int64, input models.UpdateListData) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, ListId, input)
}
