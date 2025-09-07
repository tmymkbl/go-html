package views

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var TemplateFuncs map[string]any = map[string]any{
	"set": func(renderArgs map[string]interface{}, key string, value interface{}) template.JS {
		renderArgs[key] = value
		return template.JS("")
	},
	"default": func(defVal interface{}, args ...interface{}) (interface{}, error) {
		if len(args) >= 2 {
			return nil, fmt.Errorf("wrong number of args for default: want 2 got %d", len(args)+1)
		}
		args = append(args, defVal)
		for _, val := range args {
			switch val.(type) {
			case nil:
				continue
			case string:
				if val == "" {
					continue
				}
				return val, nil
			default:
				return val, nil
			}
		}
		return nil, nil
	},
	// Replaces newlines with <br>
	"nl2br": func(text string) template.HTML {
		return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
	},
	// Skips sanitation on the parameter.  Do not use with dynamic data.
	"raw": func(text string) template.HTML {
		return template.HTML(text)
	},
	// Same as "raw"
	"unescape": func(text string) template.HTML {
		return template.HTML(text)
	},
}

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
	templates.ExecuteTemplate(writer, "layout", data)
}

func GenerateUserHTML(w http.ResponseWriter, data any, filenames ...string) {
	var files []string
	for _, file := range filenames {
		fmt.Println("Output: " + fmt.Sprintf("app/views/template_bootstrap/users/%s.html", file))
		files = append(files, fmt.Sprintf("app/views/template_bootstrap/users/%s.html", file))
	}
	templates := template.Must(template.New("todo").Funcs(TemplateFuncs).ParseFiles(files...))
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
