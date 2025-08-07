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

func generatePublicHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/public/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	// if data == nil {
	// 	templates.Execute(writer, "layout")
	// } else {
	templates.ExecuteTemplate(writer, "layout", data)
	// }
}

func generateAuthHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/auth/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	// if data == nil {
	// 	templates.Execute(writer, "layout")
	// } else {
	templates.ExecuteTemplate(writer, "layout", data)
	// }
}

func generateAdminHTML(writer http.ResponseWriter, data any, filenames ...string) {
	var files []string
	for _, file := range filenames {
		fmt.Println("Output: " + fmt.Sprintf("app/views/template_bootstrap/admin/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/admin/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "admin_layout", data)
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
	http.Handle("GET /assets/", http.StripPrefix("/assets/", files))

	adminURL := config.Config.AdminURL
	http.HandleFunc(adminURL, adminTop) //開発途中
	http.HandleFunc("/", index)
	http.HandleFunc("GET /about", about)
	http.HandleFunc("GET /contact", contact)
	http.HandleFunc("GET /blog-home", blogHome)
	http.HandleFunc("GET /blog-post", blogPost)
	http.HandleFunc("GET /faq", faq)
	http.HandleFunc("GET /portfolio-item", portfolioItem)
	http.HandleFunc("GET /portfolio-overview", portfolioOverview)
	http.HandleFunc("GET /pricing", pricing)
	http.HandleFunc("GET /signup", signup)
	http.HandleFunc("GET /login", login)
	http.HandleFunc("GET /logout", logout)
	http.HandleFunc("GET /authenticate", authenticate)
	http.HandleFunc("GET /forgot_password", forgotPassword)
	http.HandleFunc("/todos", todos)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	http.HandleFunc("GET /e401", e401)
	http.HandleFunc("GET /e404", e404)
	http.HandleFunc("GET /e500", e500)
}
