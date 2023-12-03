package services

import (
	"dana/anabul-rest-api/src/dto"
	"dana/anabul-rest-api/src/entities"
	"dana/anabul-rest-api/src/repositories"
	errors "dana/anabul-rest-api/src/templates/error"
)

type AuthService interface {
	Login(req dto.LoginFormRequestDTO) (*entities.User, error)
	Register(req dto.UserRegisterDTO) (*entities.User, error)
}
type AuthServiceImpl struct {
	repositories repositories.Repositories
}

func NewAuthService(repositories repositories.Repositories) *AuthServiceImpl {
	return &AuthServiceImpl{
		repositories: repositories,
	}
}

func (h *AuthServiceImpl) Login(req dto.LoginFormRequestDTO) (*entities.User, error) {
	user, err := h.repositories.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if req.Password != user.Password {
		return nil, errors.BadRequest("WRONG_PASSWORD")
	}
	return user, nil
}
func (h *AuthServiceImpl) Register(req dto.UserRegisterDTO) (*entities.User, error) {
	user := h.repositories.Create(req)

	return user, nil
}
