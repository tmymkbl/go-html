package controllers

import (
	"net/http"

	"go-sample-todo/app/models"
	"go-sample-todo/app/views"
)

func index(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "index", "footer")
}

func about(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "about", "footer")
}

func blogHome(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "blog-home", "footer")
}

func blogPost(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "blog-post", "footer")
}

func contact(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "contact", "footer")
}

func faq(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "faq", "footer")
}

func portfolioItem(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "portfolio-item", "footer")
}

func portfolioOverview(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "portfolio-overview", "footer")
}

func pricing(w http.ResponseWriter, r *http.Request) {
	sess, _ := models.GetSession(w, r)
	views.GeneratePublicHTML(w, sess, "layout", "header", "pricing", "footer")
}
