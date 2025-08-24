package views

import (
	"fmt"
	"net/http"
	"text/template"
)

func GeneratePublicHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/public/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func GenerateErrorHTML(writer http.ResponseWriter, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "content", nil)
}

func GenerateAuthHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/auth/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	// if data == nil {
	// 	templates.Execute(writer, "layout")
	// } else {
	templates.ExecuteTemplate(writer, "layout", data)
	// }
}

func GenerateUserHTML(w http.ResponseWriter, data any, filenames ...string) {
	var files []string
	for _, file := range filenames {
		// fmt.Println("Output: " + fmt.Sprintf("app/views/template_bootstrap/users/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/users/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "user_layout", data)
}

func GenerateAdminHTML(writer http.ResponseWriter, data any, filenames ...string) {
	var files []string
	for _, file := range filenames {
		fmt.Println("Output: " + fmt.Sprintf("app/views/template_bootstrap/admin/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/admin/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "admin_layout", data)
}
