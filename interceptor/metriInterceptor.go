package interceptor

import (
	"fmt"
	"net/http"
	"time"
)

type MetricInterceptor struct {
}

func NewMetricInterceptor() *MetricInterceptor {
	return &MetricInterceptor{}
}

func (i *MetricInterceptor) Intercept(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i.preHandle(w, r)
		start := time.Now().UnixMilli()

		h(w, r)

		end := time.Now().UnixMilli()
		fmt.Println("Request process time:", end-start)
		i.postHandle(w, r)
	})
}

func (i *MetricInterceptor) preHandle(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("metric intercept")
	return nil
}

func (i *MetricInterceptor) postHandle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
