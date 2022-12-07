package repo

import "github.com/perfectogo/template-service-with-kafka/models"

type TodoRepoInterface interface {
	CreateTodo(models.CreateUpdateTodoRequest) (models.CreateUpdateTodoResponse, error)
	GetAllTodos(models.GetTodosRequest) (models.GetAllTodosResponse, error)
	GetTodoById(models.GetAndDeleteTodoByIdRequest) (models.GetTodosRequest, error)
	UpdateTodo(models.CreateUpdateTodoRequest) (models.CreateUpdateTodoResponse, error)
	DeleteTodo(models.GetAndDeleteTodoByIdRequest) (models.DeleteTodoResponse, error)
	GetDeletedTodos(models.GetTodosRequest) (models.GetDeletedTodosResponse, error)
}
