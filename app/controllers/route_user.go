package controllers

import (
	"go-sample-todo/app/models"
	"go-sample-todo/app/views"
	"net/http"
)

func userBlogPost(w http.ResponseWriter, r *http.Request) {
	sess, err := models.GetSession(w, r)
	if err != nil {
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
