package controllers

import (
	"fmt"
	"log"
	"net/http"

	"go-sample-todo/app/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	// indexData := map[string]string{}
	// sess, err := session(w, r)
	sess, _ := session(w, r)
	// if err == nil {
	// 	indexData["login"] = sess.Email
	// }
	if r.RequestURI == "/" {
		generatePublicHTML(w, sess, "layout", "header", "index", "footer")
	} else {
		generateErrorHTML(w, "404")
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "about", "footer")
}

func blogHome(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "blog-home", "footer")
}

func blogPost(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "blog-post", "footer")
}

func contact(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "contact", "footer")
}

func faq(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "faq", "footer")
}

func portfolioItem(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "portfolio-item", "footer")
}

func portfolioOverview(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "portfolio-overview", "footer")
}

func pricing(w http.ResponseWriter, r *http.Request) {
	generatePublicHTML(w, nil, "layout", "header", "pricing", "footer")
}

func todos(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generatePublicHTML(w, user, "layout", "private_navbar", "todo")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		generatePublicHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)

		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)

		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)

		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(t)
		generatePublicHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(t)
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 200)
	}
}
