package interceptor

import (
	"net/http"
)

type HttpInterceptor interface {
	Intercept(h http.HandlerFunc) http.HandlerFunc
	preHandle(w http.ResponseWriter, r *http.Request) error
	postHandle(w http.ResponseWriter, r *http.Request) error
}

func Intercept(targetHandler http.HandlerFunc, interceptors ...HttpInterceptor) http.HandlerFunc {
	totalMiddleware := len(interceptors)
	if totalMiddleware == 0 {
		return targetHandler
	} else if totalMiddleware == 1 {
		return interceptors[0].Intercept(targetHandler)
	} else {
		chainedInterceptor := interceptors[totalMiddleware-1].Intercept(targetHandler)

		for i := totalMiddleware - 2; i >= 0; i-- {
			chainedInterceptor = interceptors[i].Intercept(chainedInterceptor)
		}

		return chainedInterceptor
	}

}
