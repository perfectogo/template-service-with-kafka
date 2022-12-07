package repo

type TodoRepoInterface interface {
	CreateTodo()
	GetAllTodos()
	GetTodoById()
	UpdateTodo()
	DeleteTodo()
}
