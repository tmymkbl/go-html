package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"go-sample-todo/config"
)

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

func SetRoute() http.Handler {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Config.Assets))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", files))
	mux.HandleFunc("GET "+config.Config.AdminURL, adminTop) //開発途中
	mux.HandleFunc("GET /about", about)
	mux.HandleFunc("GET /contact", contact)
	mux.HandleFunc("GET /blog-home", blogHome)
	mux.HandleFunc("GET /blog-post", blogPost)
	mux.HandleFunc("GET /faq", faq)
	mux.HandleFunc("GET /portfolio-item", portfolioItem)
	mux.HandleFunc("GET /portfolio-overview", portfolioOverview)
	mux.HandleFunc("GET /pricing", pricing)
	mux.HandleFunc("GET /signup", signup)
	mux.HandleFunc("POST /signup", signupAuth)
	mux.HandleFunc("GET /login", login)
	mux.HandleFunc("POST /login", loginAuth)
	mux.HandleFunc("GET /forgot_password", forgotPassword)

	mux.HandleFunc("/user/logout", logout)
	mux.HandleFunc("GET /user/profile", userProfile)
	mux.HandleFunc("GET /user/todos", todos)
	mux.HandleFunc("/user/todos/new", todoNew)
	mux.HandleFunc("/user/todos/save", todoSave)
	mux.HandleFunc("/user/todos/edit/", parseURL(todoEdit))
	mux.HandleFunc("/user/todos/update/", parseURL(todoUpdate))
	mux.HandleFunc("/user/todos/delete/", parseURL(todoDelete))
	mux.HandleFunc("GET /user/blog-post", userBlogPost)

	mux.HandleFunc("/", index) // トップページ及び不明URLの処理

	return mux
}
