package todo

import (
	"gotth/common"
	"net/http"
)

type TodoController struct {
	common.Controller
}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (t *TodoController) SetupRoute(router *http.ServeMux) {
}
