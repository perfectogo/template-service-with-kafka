package models

import (
	"time"
)

type Todo struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type CreateUpdateTodoRequest struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type CreateUpdateTodoResponse struct {
	Id string `json:"id"`
}

type GetDeleteTodoByIdRequest struct {
	Id string `json:"id"`
}

type GetTodoByIdResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Priority  string    `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GeleteTodoResponse struct {
	Id string `json:"id"`
}

type GetTodosRequest struct {
	Search string `json:"search"`
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
}

type GetAllTodosResponse struct {
	Todos []*GetTodoByIdResponse `json:"todos"`
	Count int                    `json:"count"`
}

type GetDeletedTodosResponse struct {
	DeletedTodos []*Todo `json:"deleted_todos"`
	Count        int     `json:"count"`
}
