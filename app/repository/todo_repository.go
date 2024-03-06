package repository

import (
	"github.com/Pugpaprika21/go-fiber/db"
	"github.com/Pugpaprika21/go-fiber/dto"
	"github.com/Pugpaprika21/go-fiber/models"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	Create(userIpID string, body dto.TodoBodyRequest) (uint, error)
	Get() []*dto.TodoQueryRow
	GetByID(id uint) *dto.TodoQueryRow
	Update(id uint, body dto.TodoBodyRequest) (uint, error)
	Delete(id uint) (bool, error)
}

type todoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository() *todoRepository {
	return &todoRepository{
		DB: db.Conn,
	}
}

func (t *todoRepository) Create(userIpID string, body dto.TodoBodyRequest) (uint, error) {
	todoResult := models.Todo{
		UserIpID:   userIpID,
		TodoText:   body.TodoText,
		TodoStatus: "pending",
	}

	result := t.DB.Model(&models.Todo{}).Create(&todoResult)
	if result.Error != nil {
		return 0, result.Error
	}
	return todoResult.ID, nil
}

func (t *todoRepository) Get() []*dto.TodoQueryRow {
	var todos []*dto.TodoQueryRow
	t.DB.Model(&models.Todo{}).Order("id DESC").Find(&todos)
	return todos
}

func (t *todoRepository) GetByID(id uint) *dto.TodoQueryRow {
	var todo *dto.TodoQueryRow
	t.DB.Model(&models.Todo{}).Where("id = ?", id).Order("id DESC").Find(&todo)
	return todo
}

func (t *todoRepository) Update(id uint, body dto.TodoBodyRequest) (uint, error) {
	todoResult := models.Todo{
		TodoText:   body.TodoText,
		TodoStatus: "success",
	}

	result := t.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(&todoResult)
	if result.Error != nil {
		return 0, result.Error
	}
	return todoResult.ID, nil
}

func (t *todoRepository) Delete(id uint) (bool, error) {
	result := t.DB.Scopes().Delete(&models.Todo{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
