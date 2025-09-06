package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo/backend/service"
)

type TodoController struct {
	todoService service.TodoService
}

// NewTodoController 控制器
func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{todoService: todoService}
}

func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//mux.Vars()
}

// GetTodoByID 更具ID获取待办事项
func (c *TodoController) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// 没参数是坏请求
		http.Error(w, "invalid todo ID", http.StatusBadRequest)
		return
	}
	todo, err := c.todoService.GetTodoByID(id)
	if err != nil {
		// 没查到数据是内部回话错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(todo)
}

// GetAllTodos 获取所有待办事项
func (c *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	todos, err := c.todoService.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
	}
	json.NewEncoder(w).Encode(todos)
}
