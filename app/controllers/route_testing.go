package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func testSliceByte(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "■string: "+"test")

	a := []byte{0x68, 0x6f, 0x67, 0x65}
	fmt.Fprintln(w, "■byte to string: "+string(a))

	b := []byte("foo")
	fmt.Fprintln(w, "■byte: ")
	fmt.Fprintln(w, b)
	fmt.Fprintln(w, "■byte to string: "+string(b))

	c := []string{"hoge", "bar", "baz"}
	fmt.Fprintln(w, "■string list: ")
	fmt.Fprintln(w, c)
}

func testJson(w http.ResponseWriter, r *http.Request) {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// レスポンスヘッダーにContent-Typeを設定
	w.Header().Set("Content-Type", "application/json")

	// サンプルデータの作成
	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
	}

	// JSONエンコードしてレスポンスとして送信
	json.NewEncoder(w).Encode(users)
}
