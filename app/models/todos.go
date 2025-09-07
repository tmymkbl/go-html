package models

import (
	"go-sample-todo/app/logs"
	"go-sample-todo/utils"
	"time"
)

type Todo struct {
	ID          int
	Title       string
	Content     string
	UserID      int
	CompletedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewTodo() *Todo {
	return &Todo{}
}

func (todo *Todo) CreateTodo(ID int) (err error) {
	cmd := `insert into todos (
		title, 
		content, 
		user_id, 
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, todo.Title, todo.Content, ID, time.Now())
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

func (todo *Todo) GetTodo(user_id int) (err error) {
	cmd := `select id, title, content, user_id from todos
	where id = ? and user_id = ? and deleted_at is null`
	// todo = Todo{}
	err = Db.QueryRow(cmd, todo.ID, user_id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.UserID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}

	return err
}

func (todo *Todo) GetTodoWithCompleted(user_id int) (err error) {
	cmd := `select id, title, content, user_id, completed_at, created_at from todos
	where id = ? and user_id = ? and deleted_at is null and completed_at is not null`
	// todo = Todo{}
	err = Db.QueryRow(cmd, todo.ID, user_id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.UserID,
		&todo.CompletedAt,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.DeletedAt)

	return err
}

func (t *Todo) GetTodos(user_id int) (todos []Todo, err error) {
	cmd := `select id, title, content, user_id, completed_at, created_at from todos where user_id = ? and deleted_at is null and completed_at is null order by created_at desc`
	rows, err := Db.Query(cmd, user_id)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.UserID,
			&todo.CompletedAt,
			&todo.CreatedAt)
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set title = ?, content = ?, updated_at = ? where id = ?`
	_, err = Db.Exec(cmd, t.Title, t.Content, time.Now(), t.ID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

func (t *Todo) CompleteTodo() (err error) {
	cmd := `update todos set completed_at = ? where id = ? and user_id = ?`
	_, err = Db.Exec(cmd, time.Now(), t.ID, t.UserID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}

func (t *Todo) DeleteTodo() (err error) {
	cmd := `update todos set deleted_at = ? where id = ? and user_id = ?`
	_, err = Db.Exec(cmd, time.Now(), t.ID, t.UserID)
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Error("Connect DB ", "error", err, "database", "golang", "funcName", funcName, "file", file, "line", line)
	}
	return err
}
