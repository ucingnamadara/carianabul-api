package handlers

import (
	"dana/anabul-rest-api/src/templates/response"
	"net/http"
)

func (h *RouterHandlerImpl) FindById(res http.ResponseWriter, req *http.Request) {
	userId := req.Header.Get("id")

	user, err := h.userService.FindById(userId)
	if err != nil {
		response.Error(res, err)
		return
	}
	response.Success(res, user)
}

func (h *RouterHandlerImpl) FindList(res http.ResponseWriter, req *http.Request) {
	user := h.userService.FindList()
	response.Success(res, user)
}
