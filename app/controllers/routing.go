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

func generatePublicHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/public/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func generateErrorHTML(writer http.ResponseWriter, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "content", nil)
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

func generateUserHTML(w http.ResponseWriter, data any, filenames ...string) {
	var files []string
	for _, file := range filenames {
		fmt.Println("Output: " + fmt.Sprintf("app/views/template_bootstrap/users/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/users/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "user_layout", data)
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

// func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
func session(_ http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("invalid session")
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

func SetRoute() {
	// files := http.FileServer(http.Dir(config.Config.Static))
	// http.Handle("/static/", http.StripPrefix("/static/", files))
	files := http.FileServer(http.Dir(config.Config.Assets))
	http.Handle("GET /assets/", http.StripPrefix("/assets/", files))

	adminURL := config.Config.AdminURL
	http.HandleFunc(adminURL, adminTop) //開発途中
	http.HandleFunc("GET /about", about)
	http.HandleFunc("GET /contact", contact)
	http.HandleFunc("GET /blog-home", blogHome)
	http.HandleFunc("GET /blog-post", blogPost)
	http.HandleFunc("GET /faq", faq)
	http.HandleFunc("GET /portfolio-item", portfolioItem)
	http.HandleFunc("GET /portfolio-overview", portfolioOverview)
	http.HandleFunc("GET /pricing", pricing)
	http.HandleFunc("GET /signup", signup)
	http.HandleFunc("POST /signup", signupAuth)
	http.HandleFunc("GET /login", login)
	http.HandleFunc("POST /login", loginAuth)
	http.HandleFunc("GET /forgot_password", forgotPassword)

	http.HandleFunc("/user/logout", logout)
	http.HandleFunc("GET /user/profile", userProfile)
	http.HandleFunc("GET /user/todos", todos)
	http.HandleFunc("/user/todos/new", todoNew)
	http.HandleFunc("/user/todos/save", todoSave)
	http.HandleFunc("/user/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/user/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/user/todos/delete/", parseURL(todoDelete))
	http.HandleFunc("GET /user/blog-post", userBlogPost)

	http.HandleFunc("/", index) // トップページ及び不明URLの処理

	http.HandleFunc("/test", test) // 軽いテスト用の実装
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
	a := []byte{0x68, 0x6f, 0x67, 0x65}
	fmt.Fprintln(w, string(a))
	b := []byte("foo")
	fmt.Fprintln(w, b)
	fmt.Fprintln(w, string(b))
	c := []string{"foo", "bar", "baz"}
	fmt.Fprintln(w, c)
}
