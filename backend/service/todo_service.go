package service

import (
	"todo/backend/model"
	"todo/backend/repository"
)

// TodoService 类似于java的接口层级 mapper
type TodoService interface {
	GetAllTodos() ([]model.Todo, error)
	GetTodoByID(id int) (*model.Todo, error)
	CreateTodo(title string) (int, error)
	UpdateTodo(id int, title string, completed bool) error
	ToggleTodoCompletion(id int, completed bool) error
	DeleteTodo(id int) error
}

/*
	简明说一下为什么todoService 会有 TodoRepository
	先看一个图：
		+---------------------+
		|     Handler/API     |  <-- 接收 HTTP 请求（如 Gin 路由）
		+---------------------+
                 ↓
		+---------------------+
		|     Service         |  <-- 业务逻辑（ todoService）
		+---------------------+
		          ↓
		+---------------------+
		|     Repository      |  <-- 数据访问（ todoRepository）
		+---------------------+
                  ↓
       	Database (MySQL/PostgreSQL/etc.)


		因为 todoService（业务逻辑层）需要调用 todoRepository（数据访问层）来操作数据库，这是典型的 分层架构 + 依赖注入 设计。
		(
			如果你有其他语言的基础，例如java，那么你可以把 Repository 看作 mapper,
			在 todoService 结构体的 repository.TodoRepository 相当于在 service 层注入 mapper
		)

*/
// todoService 待办事项业务逻辑实现
type todoService struct {
	repo repository.TodoRepository
}

// NewTodoService 创建待办事项业务逻辑层
func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}

// ToggleTodoCompletion 变更 完成动态
// 完成 -> 未完成
// 未完成 -> 完成
func (s *todoService) ToggleTodoCompletion(id int, completed bool) error {
	todo, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	todo.Completed = completed
	return s.repo.Update(todo)
}

func (s *todoService) UpdateTodo(id int, title string, completed bool) error {
	todo := &model.Todo{
		ID:        id,
		Title:     title,
		Completed: completed,
	}
	return s.repo.Update(todo)
}

func (s *todoService) CreateTodo(title string) (int, error) {
	todo := &model.Todo{
		Title:     title,
		Completed: false,
	}
	return s.repo.Create(todo)
}

func (s *todoService) GetTodoByID(id int) (*model.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *todoService) GetAllTodos() ([]model.Todo, error) {
	return s.repo.GetAll()
}
