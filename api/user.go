package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/natepboat/go-rest-api/service"
	"github.com/natepboat/go-router/contextKey"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	pathParam := r.Context().Value(contextKey.PathParam{}).(map[string]string)
	id, _ := strconv.Atoi(pathParam["id"])

	result := c.service.GetUser(id)

	response, err := json.Marshal(result)
	if err != nil {
		errors.New("parse response error")
	}

	w.Write(response)
}
