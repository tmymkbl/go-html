package controllers

import (
	"go-sample-todo/app/models"
	"go-sample-todo/app/views"
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
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "blog-post", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err == nil {
		// 構造体で渡してみる。そして構造体の中にmapと構造体を入れ子にしてみる。
		// 構造体内の項目名を小文字でつけると外部パッケージ（テンプレートなども）から参照できない。
		// 参考：https://qiita.com/Yarimizu14/items/e93097c4f4cfd5468259
		type Data2 struct {
			Value  string
			Value2 string
		}
		type Data struct {
			Sess      models.Session
			Dashboard string
			Value     map[string]any
			Value2    Data2
		}
		data2 := Data2{
			Value:  "value2-1",
			Value2: "value2-2",
		}
		map_data := map[string]any{
			"data1": "value-1",
			"data2": "value-2",
		}
		data := Data{
			Sess:      sess,
			Dashboard: "ダッシュボード",
			Value:     map_data,
			Value2:    data2,
		}
		// テンプレートでの取り出し確認
		// tmpl := "Dashboard: {{.Dashboard}}, Value.data1: {{.Value.data1}}, Value2.Value: {{.Value2.Value}}"
		// t, _ := template.New("test").Parse(tmpl)
		// fmt.Println(t.Execute(os.Stdout, data))
		views.GenerateUserHTML(w, data, "user_layout", "head", "nav", "sidenav", "main", "footer", "scripts")
	} else {
		http.Redirect(w, r, "/", 200)
	}
}
