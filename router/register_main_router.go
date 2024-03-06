package router

import (
	"github.com/Pugpaprika21/go-fiber/app/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	apiRouter := app.Group("/api")
	todoRouter := apiRouter.Group("/todo")
	todoController := controller.NewTodoController()
	todoRouter.Post("/createTodo", todoController.CreateTodo)
	todoRouter.Get("/getTodos", todoController.GetTodos)
	todoRouter.Get("/getTodo/:id", todoController.GetTodo)
	todoRouter.Put("/updateTodo/:id", todoController.UpdateTodo)
	todoRouter.Delete("/deleteTodo/:id", todoController.DeleteTodo)
}
