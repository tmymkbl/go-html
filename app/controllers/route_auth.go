package controllers

import (
	"log"
	"net/http"
	"strings"

	"go-sample-todo/app/logs"
	"go-sample-todo/app/models"
	"go-sample-todo/app/views"
	"go-sample-todo/config"
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
	Name            string `json:"name" validate:"required,min=2,max=255"`
	Email           string `json:"email" validate:"required,email,min=7,max=128"`
	Password        string `json:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password,min=8,max=32"`
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
	// retList はテンプレートに渡すデータを格納するマップ
	retList := map[string]any{
		"input": map[string]string{
			"Name":            r.PostFormValue("name"),
			"Email":           r.PostFormValue("email"),
			"Password":        r.PostFormValue("password"),
			"ConfirmPassword": r.PostFormValue("confirmpassword"),
		},
		"error": map[string]string{},
	}
	// fieldName := map[string]string{
	// 	"Name":            "氏名",
	// 	"Email":           "Eメールアドレス",
	// 	"Password":        "パスワード",
	// 	"ConfirmPassword": "パスワード確認",
	// }

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
			// value_ja := strings.Replace(value, key, fieldName[key], -1)
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
	retList, ok := validateSignupAuth(r)
	if !ok {
		views.GenerateAuthHTML(w, retList, "layout", "signup")
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
		views.GenerateAuthHTML(w, retList, "layout", "signup")
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
		views.GenerateAuthHTML(w, retList, "layout", "signup")
		return
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		views.GenerateAuthHTML(w, nil, "layout", "signup")
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
		views.GenerateAuthHTML(w, retList, "layout", "login")
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
		views.GenerateAuthHTML(w, retList, "layout", "login")
		return
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		cookie := http.Cookie{
			Name:  "_cookie",
			Value: session.UUID,
			SameSite: func() http.SameSite {
				switch strings.ToLower(config.Config.CookieSame) {
				case "lax":
					return http.SameSiteLaxMode
				case "strict":
					return http.SameSiteStrictMode
				case "none":
					return http.SameSiteNoneMode
				default:
					return http.SameSiteDefaultMode
				}
			}(),
			HttpOnly: config.Config.CookieHttps,
			Secure:   config.Config.CookieSecure, // 本番環境では true にすることを推奨
			// Expires:  session.CreatedAt.Add(time.Duration(config.Config.CookieExpire) * time.Hour),
			// Expires: session.CreatedAt.Add(8 * time.Hour),
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		views.GenerateAuthHTML(w, retList, "layout", "login")
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		views.GenerateAuthHTML(w, nil, "layout", "login")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func forgotPassword(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		views.GenerateAuthHTML(w, nil, "layout", "password")
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
