package interceptor

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AuthInterceptor struct {
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (i *AuthInterceptor) Intercept(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := i.preHandle(w, r)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("unauthorized"))
		} else {
			h(w, r)
		}
	})
}

func (i *AuthInterceptor) preHandle(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("intercept auth")

	apiKey := r.URL.Query().Get("x-api-key")
	if strings.EqualFold(apiKey, "") {
		return errors.New("no api key found")
	}

	return nil
}

func (i *AuthInterceptor) postHandle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
