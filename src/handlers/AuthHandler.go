package handlers

import (
	"dana/anabul-rest-api/src/dto"
	"dana/anabul-rest-api/src/templates/response"
	"dana/anabul-rest-api/src/utils"
	"encoding/json"
	"net/http"
)

func (h *RouterHandlerImpl) Login(res http.ResponseWriter, req *http.Request) {
	loginFormRequestDTO := dto.LoginFormRequestDTO{}
	err := json.NewDecoder(req.Body).Decode(&loginFormRequestDTO)
	if err != nil {
		response.Error(res, err)
		return
	}

	user, err := h.authService.Login(loginFormRequestDTO)
	if err != nil {
		response.Error(res, err)
		return
	}

	tokenMap, err := utils.JwtSign(user)
	if err != nil {
		response.Error(res, err)
		return
	}

	result := dto.LoginResponseDTO{}
	result.Token = tokenMap
	result.User = user
	response.Success(res, result)
}

func (h *RouterHandlerImpl) Register(res http.ResponseWriter, req *http.Request) {
	dto := dto.UserRegisterDTO{}
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		response.Error(res, err)
		return
	}

	result, err := h.authService.Register(dto)
	if err != nil {
		response.Error(res, err)
		return
	}

	response.Success(res, result)
}
