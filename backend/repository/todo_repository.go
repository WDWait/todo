package repository

import (
	"log"
	"todo/backend/config"
	"todo/backend/model"
)

// TodoRepository create todo_repository interface
type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	GetByID(id int) (*model.Todo, error)
	Create(todo *model.Todo) (int, error)
	Update(todo *model.Todo) error
	Delete(id int) error
}

// todoRepository struct
type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (t *todoRepository) Delete(id int) error {
	sqlStr := "delete from todos where id = ?"
	_, err := config.DB.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) Update(todo *model.Todo) error {
	sqlStr := "UPDATE todos SET title = ?, completed = ? WHERE id = ?"
	_, err := config.DB.Exec(sqlStr, todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *todoRepository) Create(todo *model.Todo) (int, error) {
	sqlStr := "insert into todos (title, completed) values (?,?)"
	exec, err := config.DB.Exec(sqlStr, todo.Title, todo.Completed)
	if err != nil {
		log.Panicln(err.Error())
		return 0, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		log.Fatalf("get LastInsertId err,err: %v", err.Error())
		return 0, err
	}
	return int(id), nil
}

func (t *todoRepository) GetByID(id int) (*model.Todo, error) {
	getByID := "SELECT id, title, completed, created_at, updated_at FROM todos where id = ?"

	var todo model.Todo
	err := config.DB.QueryRow(getByID, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	/*
		Tips: 简明的说
		QueryRow() 查询后不需要手动调用 Close()，是因为 QueryRow 内部已经自动处理了资源的释放，在 Scan() 之后。
		Query() 返回的是 *sql.Rows，它是一个游标（cursor），可能包含多行数据，Go 无法预知你什么时候读完，所以必须手动关闭，否则会导致数据库连接泄漏。
	*/
	return &todo, nil
}

func (t *todoRepository) GetAll() ([]model.Todo, error) {
	getAllSql := "SELECT id, title, completed, created_at, updated_at FROM todos ORDER BY created_at DESC"
	rows, err := config.DB.Query(getAllSql)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// very important : close rows and release the held database link
	defer rows.Close()

	// do map
	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
