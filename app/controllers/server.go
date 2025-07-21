package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"go-sample-todo/app/models"
	"go-sample-todo/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		fmt.Println("Output: " + file + ".html")
		// files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	// templates.ExecuteTemplate(writer, "layout", data)
	templates.Execute(writer, data)
}

func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		id, _ := strconv.Atoi(q[2])
		fmt.Println(id)
		fn(w, r, id)
	}
}

func StartMainServer() {
	// files := http.FileServer(http.Dir(config.Config.Static))
	// http.Handle("/static/", http.StripPrefix("/static/", files))
	files := http.FileServer(http.Dir(config.Config.Assets))
	http.Handle("/assets/", http.StripPrefix("/assets/", files))

	http.HandleFunc("/", top) //top
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	return
	// return http.ListenAndServe(":"+config.Config.Port, nil)
}
