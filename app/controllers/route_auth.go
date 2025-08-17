package controllers

import (
	"log"
	"net/http"
	"strings"

	"go-sample-todo/app/logs"
	"go-sample-todo/app/models"
	"go-sample-todo/utils"

	en "github.com/go-playground/locales/en"
	ja "github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

type LoginUser struct {
	Email    string `validate:"required,min=7,max=128,email"`
	Password string `validate:"required,min=8,max=32"`
}

type SignupUser struct {
	Name            string `validate:"required,max=255"`
	Email           string `validate:"required,email,min=7,max=128"`
	Password        string `validate:"required,min=8,max=32"`
	ConfirmPassword string `validate:"required,eqfield=Password,min=8,max=32"`
}

// validateAuthenticate は login form data の検証を行います。
// 項目 email と password フィールドを検証します。
// 検証に失敗した場合は、ログを出力し、false を返します。
func validateSignupAuth(r *http.Request) (map[string]any, bool) {
	uni := ut.New(en.New(), ja.New())
	trans, _ := uni.GetTranslator("ja")
	err := r.ParseForm()
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Info("signup validation error ", "error", err, "file", file, "func", funcName, "line", line)
		return map[string]any{}, false
	}
	retList := map[string]any{
		"input": map[string]string{
			"Name":            r.PostFormValue("name"),
			"Email":           r.PostFormValue("email"),
			"Password":        r.PostFormValue("password"),
			"ConfirmPassword": r.PostFormValue("confirmpassword"),
		},
		"error": map[string]string{},
	}

	validate := validator.New()
	ja_translations.RegisterDefaultTranslations(validate, trans)
	vErr := validate.Struct(
		SignupUser{
			Name:            r.PostFormValue("name"),
			Email:           r.PostFormValue("email"),
			Password:        r.PostFormValue("password"),
			ConfirmPassword: r.PostFormValue("confirmpassword"),
		})
	if vErr != nil {
		e := vErr.(validator.ValidationErrors)
		jaErrMsg := e.Translate(trans) // jaErrMsg は map[string]string 型で、キーは LoginUser.フィールド名、値は日本語のエラーメッセージです。
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Debug("signup validation error ", "errors", jaErrMsg, "file", file, "func", funcName, "line", line)
		errorMap := retList["error"].(map[string]string)
		for key, value := range jaErrMsg {
			errorMap[strings.Replace(key, ".", "_", -1)] = value
		}
		retList["error"] = errorMap
		return retList, false
	}
	file, line, funcName := utils.GetCurrentInfo()
	logs.Log.Info("signup Validation successful! ", "user", r.PostFormValue("email"), "file", file, "func", funcName, "line", line)
	return retList, true
}

func signupAuth(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	log.Println(err)
	// }
	// user := models.User{
	// 	Name:     r.PostFormValue("name"),
	// 	Email:    r.PostFormValue("email"),
	// 	PassWord: r.PostFormValue("password"),
	// }
	// if err := user.CreateUser(); err != nil {
	// 	log.Println(err)
	// }

	// http.Redirect(w, r, "/", 200)

	// retList := map[string]string{
	// 	"Email": "Eメールの入力が誤っています",
	// }
	retList, ok := validateSignupAuth(r)
	if !ok {
		generateAuthHTML(w, retList, "layout", "signup")
		return
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err == nil {
		authErrList := map[string]string{
			"main": "何らかの障害が発生しました。しばらくしてから再度お試しください。", // 登録済みのメールアドレスを知らせないため、あえてこのメッセージ
		}
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Info("signup Error ", "email", r.PostFormValue("email"), "message", err, "file", file, "func", funcName, "line", line)
		retList["error"] = authErrList
		generateAuthHTML(w, retList, "layout", "signup")
		return
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		generateAuthHTML(w, retList, "layout", "signup")
		return
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	/*
		if r.Method == "GET" {
			_, err := session(w, r)
			if err != nil {
				generateAuthHTML(w, nil, "layout", "public_navbar", "signup")
			} else {
				http.Redirect(w, r, "/todos", 302)
			}
		} else if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
			}
			user := models.User{
				Name:     r.PostFormValue("name"),
				Email:    r.PostFormValue("email"),
				PassWord: r.PostFormValue("password"),
			}
			if err := user.CreateUser(); err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/", 302)
		}
	*/

	_, err := session(w, r)
	if err != nil {
		generateAuthHTML(w, nil, "layout", "signup")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// validateAuthenticate は login form data の検証を行います。
// 項目 email と password フィールドを検証します。
// 検証に失敗した場合は、ログを出力し、false を返します。
func validateLoginAuth(r *http.Request) (map[string]any, bool) {
	uni := ut.New(en.New(), ja.New())
	trans, _ := uni.GetTranslator("ja")
	err := r.ParseForm()
	if err != nil {
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Info("login validation error ", "error", err, "file", file, "func", funcName, "line", line)
		return map[string]any{}, false
	}
	retList := map[string]any{
		"input": map[string]string{
			"Email":    r.PostFormValue("email"),
			"Password": r.PostFormValue("password"),
		},
		"error": map[string]string{},
	}
	validate := validator.New()
	ja_translations.RegisterDefaultTranslations(validate, trans)
	vErr := validate.Struct(
		LoginUser{
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		})
	if vErr != nil {
		e := vErr.(validator.ValidationErrors)
		jaErrMsg := e.Translate(trans) // jaErrMsg は map[string]string 型で、キーは LoginUser.フィールド名、値は日本語のエラーメッセージです。
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Debug("login validation error ", "errors", jaErrMsg, "file", file, "func", funcName, "line", line)
		// errorMap := retList["error"].(map[string]string)
		// for _, v := range vErr.(validator.ValidationErrors) {
		// 	v.Translate(trans)
		// 	file, line, funcName := utils.GetCurrentInfo()
		// 	logs.Log.Debug("login validation error ", "message", v.Error(), "field", v.Field(), "tag", v.Tag(), "param", v.Param(), "file", file, "func", funcName, "line", line)
		// 	// fmt.Println(errv.Field(), errv.Tag(), errv.Param())
		// 	errorMap[v.Field()] = v.Tag() + v.Param()
		// }
		// retList["error"] = errorMap
		errorMap := retList["error"].(map[string]string)
		for key, value := range jaErrMsg {
			errorMap[strings.Replace(key, ".", "_", -1)] = value
		}
		retList["error"] = errorMap
		return retList, false
	}
	file, line, funcName := utils.GetCurrentInfo()
	logs.Log.Info("login Validation successful! ", "user", r.PostFormValue("email"), "file", file, "func", funcName, "line", line)
	return retList, true
}

// authenticate handles the login process.
// It checks if the user credentials are valid and creates a session if they are.
func loginAuth(w http.ResponseWriter, r *http.Request) {
	retList, ok := validateLoginAuth(r)
	if !ok {
		generateAuthHTML(w, retList, "layout", "login")
		return
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		authErrList := map[string]string{
			"main": "Eメールとパスワードのいずれかまたは両方の入力が誤っています",
		}
		file, line, funcName := utils.GetCurrentInfo()
		logs.Log.Info("login Error ", "email", r.PostFormValue("email"), "message", err, "file", file, "func", funcName, "line", line)
		retList["error"] = authErrList
		generateAuthHTML(w, retList, "layout", "login")
		return
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		generateAuthHTML(w, retList, "layout", "login")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateAuthHTML(w, nil, "layout", "login")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func forgotPassword(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateAuthHTML(w, nil, "layout", "password")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
