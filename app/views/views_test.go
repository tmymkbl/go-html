package views

import (
	"bytes"
	"fmt"
	"go-sample-todo/app/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
