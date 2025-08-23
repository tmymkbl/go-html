package controllers

import (
	"go-sample-todo/app/models"
	"net/http"
)

func userBlogPost(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
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
		generateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "blog-post", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err == nil {
		// 配列から構造体に変更してみた
		type Data struct {
			sess      models.Session
			Dashboard string
			value     map[string]any
		}

		value_data := map[string]any{
			"data1": "value1",
			"data2": "value2",
		}
		data := Data{
			sess:      sess,
			Dashboard: "ダッシュボード",
			value:     value_data,
		}
		generateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "main", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}
