package controllers

import (
	"go-sample-todo/app/logs"
	"go-sample-todo/app/models"
	"go-sample-todo/app/views"
	"go-sample-todo/utils"
	"io"
	"net/http"
	"os"
)

func userBlogPost(w http.ResponseWriter, r *http.Request) {
	sess, err := models.GetSession(w, r)
	if err == nil {
		data := map[string]any{
			"sess":      sess,
			"Dashboard": "ダッシュボード",
			"value":     map[string]any{},
		}
		data["value"] = map[string]any{
			"data1": "value1",
			"data2": "value2",
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "blog-post", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}

func userBlogPostEntry(w http.ResponseWriter, r *http.Request) {
	sess, err := models.GetSession(w, r)
	if err == nil {

		err := r.ParseMultipartForm(32 << 20) // maxMemory
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("titleimg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		f, err := os.Create("/tmp/title.jpg")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		io.Copy(f, file)

		data := map[string]any{
			"sess":      sess,
			"Dashboard": "ダッシュボード",
			"value":     map[string]any{},
		}
		data["value"] = map[string]any{
			"data1": "value1",
			"data2": "value2",
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "blog-post", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	sess, err := models.GetSession(w, r)
	if err == nil {
		// 配列から構造体に変更してみた
		type Data struct {
			Sess      models.Session
			Dashboard string
			Value     map[string]any
		}

		value_data := map[string]any{
			"data1": "value1",
			"data2": "value2",
		}
		data := Data{
			Sess:      sess,
			Dashboard: "ダッシュボード",
			Value:     value_data,
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "main", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}

// ユーザーダッシュボード表示
func userDashboard(w http.ResponseWriter, r *http.Request) {
	sess, err := models.GetSession(w, r)
	if err == nil {
		// 配列から構造体に変更してみた
		type Data struct {
			Sess      models.Session
			Dashboard string
			Value     map[string]any
		}

		value_data := map[string]any{
			"data1": "value1",
			"data2": "value2",
		}
		data := Data{
			Sess:      sess,
			Dashboard: "ダッシュボード",
			Value:     value_data,
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "main", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}

// TODO一覧画面表示
func todos(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Sess      models.Session
		Dashboard string
		Todos     []models.Todo
	}
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, _ := models.GetUserByEmail(sess.Email)
		todos, _ := user.GetTodosByUser()
		data := Data{
			Sess:  sess,
			Todos: todos,
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "todo", "footer", "scripts")
	}
}

// TODOの新規登録画面表示
func todoNew(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Sess      models.Session
		Dashboard string
		Todos     []models.Todo
	}
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, _ := models.GetUserByEmail(sess.Email)
		todos, _ := user.GetTodosByUser()
		data := Data{
			Sess:  sess,
			Todos: todos,
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "todo_new", "footer", "scripts")
	}
}

// TODOの新規保存
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		t := models.NewTodo()
		t.Title = r.PostFormValue("title")
		t.Content = r.PostFormValue("content")
		if err := t.CreateTodo(user.ID); err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		http.Redirect(w, r, "/user/todos", http.StatusFound)
	}
}

// TODOの更新画面の表示
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	type Data struct {
		Sess      models.Session
		Dashboard string
		TodoId    int
		Todos     models.Todo
	}
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		todo := models.NewTodo()
		todo.ID = id
		err := todo.GetTodo(sess.UserID)
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		data := Data{
			Sess:   sess,
			TodoId: id,
			Todos:  *todo,
		}
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "todo_edit", "footer", "scripts")
	}
}

// TODOの更新処理
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		// TODOデータの更新
		title := r.PostFormValue("title")
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Title: title, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		// todosの画面に遷移
		todos(w, r)
	}
}

func todoComplete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		todo := models.NewTodo()
		todo.ID = id
		todo.UserID = sess.UserID
		if err := todo.CompleteTodo(); err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		todos(w, r)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := models.GetSession(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		todo := models.NewTodo()
		todo.ID = id
		todo.UserID = sess.UserID
		if err := todo.DeleteTodo(); err != nil {
			file, line, funcName := utils.GetCurrentInfo()
			logs.Log.Error("main ConnectDB ", "error", err, "file", file, "line", line, "funcName", funcName)
		}
		todos(w, r)
	}
}
