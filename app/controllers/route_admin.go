package controllers

import (
	"net/http"
)

func adminTop(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	data := map[string]any{
		"sess":      sess,
		"Dashboard": "ダッシュボード",
		"value":     map[string]any{},
	}
	data["value"] = map[string]any{
		"data1": "value1",
		"data2": "value2",
	}
	if err != nil {
		generateAdminHTML(w, data, "admin_layout", "head", "nav", "sidenav", "main", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
