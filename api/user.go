package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/natepboat/go-rest-api/provider"
	"github.com/natepboat/go-rest-api/service"
	"github.com/natepboat/go-router/contextKey"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(provider *provider.ComponentProvider) *UserController {
	return &UserController{
		userService: provider.Required("service.UserService").(service.IUserService),
	}
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	pathParam := r.Context().Value(contextKey.PathParam{}).(map[string]string)
	id, _ := strconv.Atoi(pathParam["id"])

	result := c.userService.GetUser(id)

	response, err := json.Marshal(result)
	if err != nil {
		errors.New("parse response error")
	}

	w.Write(response)
}
