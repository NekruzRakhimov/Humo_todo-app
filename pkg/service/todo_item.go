package service

import (
	"todo-app/models"
	"todo-app/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int64, item models.TodoItem) (int64, error) {
	_, err := s.listRepo.GetById(int(userId), int(listId))
	if err != nil {
		//list doesn't exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int64) ([]models.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int64) (models.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(listId, itemId int64) error {
	return s.repo.Delete(listId, itemId)
}

func (s *TodoItemService) Update(ListId int64, itemId int64, input models.UpdateItemData) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ListId, itemId, input)
}