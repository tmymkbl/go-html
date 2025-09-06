package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"go-sample-todo/config"
)

var todoPath = regexp.MustCompile("^/user/todos/(edit|save|update|complete|delete)/([0-9]+)$")

func parseURL(reg regexp.Regexp, fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		myUrl, _ := url.Parse(r.URL.String())
		fmt.Println(myUrl)
		q := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		id, _ := strconv.Atoi(q[0][2])
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
	mux.HandleFunc("GET /user", userDashboard)
	mux.HandleFunc("GET /user/profile", userProfile)
	mux.HandleFunc("GET /user/todos", todos)
	mux.HandleFunc("GET /user/todos/new", todoNew)
	mux.HandleFunc("POST /user/todos/save", todoSave)
	mux.HandleFunc("GET /user/todos/edit/", parseURL(*todoPath, todoEdit))
	mux.HandleFunc("POST /user/todos/update/", parseURL(*todoPath, todoUpdate))
	mux.HandleFunc("GET /user/todos/complete/", parseURL(*todoPath, todoComplete))
	mux.HandleFunc("GET /user/todos/delete/", parseURL(*todoPath, todoDelete))
	mux.HandleFunc("GET /user/blog-post", userBlogPost)
	mux.HandleFunc("POST /user/blog-post", userBlogPostEntry)

	mux.HandleFunc("/", index) // トップページ及び不明URLの処理

	return mux
}
