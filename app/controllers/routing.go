package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"go-sample-todo/app/models"
	"go-sample-todo/config"
)

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

	http.HandleFunc("GET "+config.Config.AdminURL, adminTop) //開発途中
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

	http.HandleFunc("GET /trial/slicebyte", trialSliceByte) // スライステスト用の実装
	http.HandleFunc("GET /trial/json", trialJson)           // jsonテスト用の実装

	http.HandleFunc("/", index) // トップページ及び不明URLの処理
}
