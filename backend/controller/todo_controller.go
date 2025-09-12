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

// DeleteTodo 删除待办项目
func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// 没有参数响应坏请求
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = c.todoService.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateTodo 更新
func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		return
	}

	var request struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = c.todoService.UpdateTodo(id, request.Title, request.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// StatusNoContent 无响应内容
	w.WriteHeader(http.StatusNoContent)
}

// CreateTodo 穿件待办事项
func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var request struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := c.todoService.CreateTodo(request.Title)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// ToggleTodoCompletion 切换待办事项的完成状态
func (c *TodoController) ToggleTodoCompletion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var request struct {
		Completed bool `json:"completed"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.todoService.ToggleTodoCompletion(id, request.Completed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
