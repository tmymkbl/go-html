package controllers

import (
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
		data := map[string]any{
			"sess":      sess,
			"Dashboard": "ダッシュボード",
			"value":     map[string]any{},
		}
		data["value"] = map[string]any{
			"data1": "value1",
			"data2": "value2",
		}
		generateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "main", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}
