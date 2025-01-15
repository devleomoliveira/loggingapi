package common

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(w http.ResponseWriter, code int, data interface{}, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	NewResponse(w, http.StatusOK, data, "success")
}

func ResponseFailed(w http.ResponseWriter, r *http.Request, code int, err error) {
	if code == 0 {
		code = http.StatusInternalServerError
	}

	if code == http.StatusUnauthorized {
		if cookie, err := r.Cookie(CookieTokenName); err == nil && cookie.Value != "" {
			http.SetCookie(w, &http.Cookie{Name: CookieTokenName, Value: "", MaxAge: -1, Path: "/", Secure: true, HttpOnly: true})
			http.SetCookie(w, &http.Cookie{Name: CookieLoginUser, Value: "", MaxAge: -1, Path: "/", Secure: true})
		}
	}

	var message string
	if err != nil {
		message = err.Error()
		user := GetUserFromRequest(r)

		var url string
		if r != nil {
			url = r.URL.String()
		}
		logrus.Warnf("url: %s, user: %s, error: %v", url, user.Name, message)
	}
	NewResponse(w, code, nil, message)
}

// TODO: trocar por func em context
func GetUserFromRequest(_ *http.Request) struct {
	Name string
} {
	result := struct {
		Name string
	}{
		Name: "Teste",
	}
	return result
}
