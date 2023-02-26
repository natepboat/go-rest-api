package api

import (
	"net/http"

	"github.com/natepboat/go-router/contextKey"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	pathParam := r.Context().Value(contextKey.PathParam{}).(map[string]string)
	w.Write([]byte("getuser: " + pathParam["id"]))
}
