package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Pugpaprika21/go-fiber/app/repository"
	"github.com/Pugpaprika21/go-fiber/dto"
	"github.com/gofiber/fiber/v2"
)

type todoController struct {
	todoRepository repository.ITodoRepository
	validate       repository.IValidatorRepository
}

func NewTodoController() *todoController {
	return &todoController{
		todoRepository: repository.NewTodoRepository(),
		validate:       repository.NewXValidator(),
	}
}

func (t *todoController) CreateTodo(c *fiber.Ctx) error {
	var body dto.TodoBodyRequest
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errs := t.validate.Validator(&body)
	if len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	todoID, err := t.todoRepository.Create(c.IP(), body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": todoID})
}

func (t *todoController) GetTodos(c *fiber.Ctx) error {
	var todosResp []dto.TodoRespone
	todos := t.todoRepository.Get()
	for _, todo := range todos {
		todosResp = append(todosResp, dto.TodoRespone{
			ID:           todo.ID,
			CreatedAt:    todo.CreatedAt,
			UpdatedAt:    todo.UpdatedAt,
			DeletedAt:    todo.DeletedAt,
			UserIpID:     todo.UserIpID,
			TodoStatus:   todo.TodoStatus,
			TodoText:     todo.TodoText,
			ActiveStatus: todo.ActiveStatus,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": todosResp})
}

func (t *todoController) GetTodo(c *fiber.Ctx) error {
	todoID, _ := strconv.ParseUint(c.Params("id"), 0, 0)
	todo := t.todoRepository.GetByID(uint(todoID))
	todoResp := dto.TodoRespone{
		ID:           todo.ID,
		CreatedAt:    todo.CreatedAt,
		UpdatedAt:    todo.UpdatedAt,
		DeletedAt:    todo.DeletedAt,
		UserIpID:     todo.UserIpID,
		TodoStatus:   todo.TodoStatus,
		TodoText:     todo.TodoText,
		ActiveStatus: todo.ActiveStatus,
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": todoResp})
}

func (t *todoController) UpdateTodo(c *fiber.Ctx) error {
	todoID, _ := strconv.ParseUint(c.Params("id"), 0, 0)

	var body dto.TodoBodyRequest
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	errs := t.validate.Validator(&body)
	if len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	_, err := t.todoRepository.Update(uint(todoID), body)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": todoID, "body": body})
}

func (t *todoController) DeleteTodo(c *fiber.Ctx) error {
	todoID, _ := strconv.ParseUint(c.Params("id"), 0, 0)

	_, err := t.todoRepository.Delete(uint(todoID))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": nil})
}
