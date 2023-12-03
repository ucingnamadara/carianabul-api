package handlers

import (
	"dana/anabul-rest-api/src/middlewares"
	"dana/anabul-rest-api/src/services"

	"github.com/gorilla/mux"
)

type RouterHandlerImpl struct {
	userService services.UserService
	authService services.AuthService
}

func NewHttpHandler(userService services.UserService, authService services.AuthService) *RouterHandlerImpl {
	return &RouterHandlerImpl{
		userService: userService,
		authService: authService,
	}
}
func (h *RouterHandlerImpl) Router(r *mux.Router) {
	r.HandleFunc("/user/me", middlewares.VerifyJWT(h.FindById)).Methods("GET", "OPTIONS")

	r.HandleFunc("/auth/login", h.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/auth/register", h.Register).Methods("POST", "OPTIONS")
}
