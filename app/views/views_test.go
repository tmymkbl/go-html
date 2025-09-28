package views

import (
	"bytes"
	"fmt"
	"go-sample-todo/app/models"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTemplateFuncs_set(t *testing.T) {
	fn, ok := TemplateFuncs["set"].(func(map[string]interface{}, string, interface{}) template.JS)
	if !ok {
		t.Fatalf("TemplateFuncs[\"set\"] has wrong type")
	}
	m := map[string]interface{}{}
	ret := fn(m, "foo", 123)
	if m["foo"] != 123 {
		t.Errorf("set did not set value, got %v", m["foo"])
	}
	if ret != template.JS("") {
		t.Errorf("set did not return empty template.JS, got %v", ret)
	}
}

func TestTemplateFuncs_default(t *testing.T) {
	fn, ok := TemplateFuncs["default"].(func(interface{}, ...interface{}) (interface{}, error))
	if !ok {
		t.Fatalf("TemplateFuncs[\"default\"] has wrong type")
	}
	tests := []struct {
		name string
		def  interface{}
		args []interface{}
		want interface{}
	}{
		{"nil arg", "def", []interface{}{nil}, "def"},
		{"empty string", "def", []interface{}{""}, "def"},
		{"non-empty string", "def", []interface{}{"abc"}, "abc"},
		{"non-string", "def", []interface{}{int(42)}, 42},
		{"no args", "", []interface{}{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fn(tt.def, tt.args...)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("default(%v, %v) = %v, want %v", tt.def, tt.args, got, tt.want)
			}
		})
	}
	// Too many args
	_, err := fn("def", 1, 2)
	if err == nil {
		t.Errorf("expected error for too many args")
	}
}

func TestTemplateFuncs_nl2br(t *testing.T) {
	fn, ok := TemplateFuncs["nl2br"].(func(string) template.HTML)
	if !ok {
		t.Fatalf("TemplateFuncs[\"nl2br\"] has wrong type")
	}
	input := "a\n<b>&"
	want := template.HTML("a<br>&lt;b&gt;&amp;")
	got := fn(input)
	if got != want {
		t.Errorf("nl2br(%q) = %q, want %q", input, got, want)
	}
}

func TestTemplateFuncs_raw(t *testing.T) {
	fn, ok := TemplateFuncs["raw"].(func(string) template.HTML)
	if !ok {
		t.Fatalf("TemplateFuncs[\"raw\"] has wrong type")
	}
	input := "<b>raw</b>"
	got := fn(input)
	if got != template.HTML(input) {
		t.Errorf("raw(%q) = %q, want %q", input, got, input)
	}
}

func TestTemplateFuncs_unescape(t *testing.T) {
	fn, ok := TemplateFuncs["unescape"].(func(string) template.HTML)
	if !ok {
		t.Fatalf("TemplateFuncs[\"unescape\"] has wrong type")
	}
	input := "<b>unescape</b>"
	got := fn(input)
	if got != template.HTML(input) {
		t.Errorf("unescape(%q) = %q, want %q", input, got, input)
	}
}

func TestGenerateErrorHTML(t *testing.T) {
	// httptest.ResponseRecorderでレスポンスを受け取る
	rec := httptest.NewRecorder()

	// 存在するエラーテンプレートファイル名を指定（例: "404" や "error" などプロジェクトに合わせて変更）
	// テンプレートファイルが存在しない場合はエラーになるので注意
	GenerateErrorHTML(rec, "404")

	// ステータスコードの確認（template.ExecuteTemplateはデフォルトで200を返す）
	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but got %d", http.StatusOK, rec.Code)
	}

	// レスポンスボディの内容を確認
	body := rec.Body.String()
	if len(body) == 0 {
		t.Errorf("response body is empty")
	}
	// 例: "404" テンプレートに特定の文字列が含まれていることを確認したい場合
	// if !strings.Contains(body, "Not Found") {
	//     t.Errorf("response body does not contain expected content: %s", body)
	// }
}

func TestGenerateUserHTML(t *testing.T) {
	// Arrange
	// http.Requestオブジェクトを生成
	// reqBody := bytes.NewBufferString("request body")
	// req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// 生成後は*http.Requestオブジェクトと同じように扱える
	//   q := req.URL.Query()
	//   q.Add("", tt.args.code)
	//   req.URL.RawQuery = q.Encode()

	// レスポンスを受け取る*httptest.ResponseRecorder
	got := httptest.NewRecorder()

	// Act
	type Data struct {
		Sess      models.Session
		Dashboard string
		TodoId    int
		Todos     models.Todo
	}
	id := 1
	data := Data{
		TodoId: id,
	}
	GenerateUserHTML(got, data, "user_layout", "head", "nav", "sidenav", "todo_edit", "footer", "scripts")

	// Assertion
	// http.Clientなどで受け取ったhttp.Responseを検証するときとほぼ変わらない
	if got.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, got.Code)
	}
	// Bodyは*bytes.Buffer型なので文字列に変換して比較 "<html"と"/html>"が入っていることを確認してみた。
	body := got.Body.String()
	if strings.Contains("<html", body) && strings.Contains("/html>", body) {
		t.Errorf("want %s, but nothing\nHTML:\n%s", "<html", body)
	}
	// Bodyの文字数で比較
	ret := len(body)
	if ret < 1 {
		t.Errorf("got.Result().ContentLength was 0")
		fmt.Println(ret)
	}
	// http.Responseオブジェクトを受け取って比較もできるみたい。
	// if resp := got.Result(); resp.ContentLength == 0 {
	// 	t.Errorf("resp.ContentLength was 0")
	// }
	// fmt.Println("got.Result() ------------------")
	// fmt.Println(got.Result())
	// fmt.Println("got.Code ------------------")
	// fmt.Println(got.Code)
	// fmt.Println("got.Header ----------------")
	// fmt.Println(got.Header())
	// fmt.Println("got.Body -----------------")
	// fmt.Println(got.Body)
	// fmt.Println(got.Result().StatusCode)
	// fmt.Println(got.Flushed)
}

func TestGeneratePublicHTML(t *testing.T) {
	// Arrange
	// http.Requestオブジェクトを生成
	reqBody := bytes.NewBufferString("request body")
	req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// 生成後は*http.Requestオブジェクトと同じように扱える
	//   q := req.URL.Query()
	//   q.Add("", tt.args.code)
	//   req.URL.RawQuery = q.Encode()

	// レスポンスを受け取る*httptest.ResponseRecorder
	got := httptest.NewRecorder()

	// Act
	sess, _ := models.GetSession(got, req)
	GeneratePublicHTML(got, sess, "layout", "header", "index", "footer")

	// Assertion
	// http.Clientなどで受け取ったhttp.Responseを検証するときとほぼ変わらない
	if got.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, got.Code)
	}
	// Bodyは*bytes.Buffer型なので文字列に変換して比較 "<html"が入っていることを確認してみた。
	body := got.Body.String()
	if strings.Contains("<html", body) {
		t.Errorf("want %s, but nothing\nHTML:\n%s", "<html", body)
	}
	// Bodyの文字数で比較
	ret := len(body)
	if ret < 1 {
		t.Errorf("got.Result().ContentLength was 0")
		fmt.Println(ret)
	}
}

func TestGenerateAuthHTML(t *testing.T) {
	rec := httptest.NewRecorder()

	// テスト用のデータを用意（必要に応じて構造体などに変更してください）
	data := map[string]interface{}{
		"Message": "Test Auth",
	}

	// 存在するauthテンプレートファイル名を指定（例: "layout", "login" などプロジェクトに合わせて変更）
	GenerateAuthHTML(rec, data, "layout", "login")

	// ステータスコードの確認
	if rec.Code != http.StatusOK {
		t.Errorf("want %d, but got %d", http.StatusOK, rec.Code)
	}

	// レスポンスボディの内容を確認
	body := rec.Body.String()
	if len(body) == 0 {
		t.Errorf("response body is empty")
	}
	// 必要に応じて、特定の文字列が含まれているかチェック
	// if !strings.Contains(body, "ログイン") {
	//     t.Errorf("response body does not contain expected content: %s", body)
	// }
}

func TestGenerateAdminHTML(t *testing.T) {
	// Arrange
	// http.Requestオブジェクトを生成
	// reqBody := bytes.NewBufferString("request body")
	// req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/user", reqBody)

	// 生成後は*http.Requestオブジェクトと同じように扱える
	//   q := req.URL.Query()
	//   q.Add("", tt.args.code)
	//   req.URL.RawQuery = q.Encode()

	// レスポンスを受け取る*httptest.ResponseRecorder
	got := httptest.NewRecorder()

	// Act
	type Data struct {
		Sess      models.Session
		Dashboard string
		TodoId    int
		Todos     models.Todo
	}
	id := 1
	data := Data{
		TodoId: id,
	}
	GenerateAdminHTML(got, data, "admin_layout", "head", "nav", "sidenav", "main", "footer", "scripts")

	// Assertion
	// http.Clientなどで受け取ったhttp.Responseを検証するときとほぼ変わらない
	if got.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, got.Code)
	}
	// Bodyは*bytes.Buffer型なので文字列に変換して比較 "<html"と"/html>"が入っていることを確認してみた。
	body := got.Body.String()
	if strings.Contains("<html", body) && strings.Contains("/html>", body) {
		t.Errorf("want %s, but nothing\nHTML:\n%s", "<html", body)
	}
	// Bodyの文字数で比較
	ret := len(body)
	if ret < 1 {
		t.Errorf("got.Result().ContentLength was 0")
		fmt.Println(ret)
	}
	// http.Responseオブジェクトを受け取って比較もできるみたい。
	// if resp := got.Result(); resp.ContentLength == 0 {
	// 	t.Errorf("resp.ContentLength was 0")
	// }
	// fmt.Println("got.Result() ------------------")
	// fmt.Println(got.Result())
	// fmt.Println("got.Code ------------------")
	// fmt.Println(got.Code)
	// fmt.Println("got.Header ----------------")
	// fmt.Println(got.Header())
	// fmt.Println("got.Body -----------------")
	// fmt.Println(got.Body)
	// fmt.Println(got.Result().StatusCode)
	// fmt.Println(got.Flushed)
}
