package services

import (
	"dana/anabul-rest-api/src/entities"
	"dana/anabul-rest-api/src/repositories"
)

type UserService interface {
	FindById(userId string) (*entities.User, error)
	FindList() *entities.UserList
}

type UserServiceImpl struct {
	repositories repositories.Repositories
}

func NewUserService(repositories repositories.Repositories) *UserServiceImpl {
	return &UserServiceImpl{
		repositories: repositories,
	}
}

func (h *UserServiceImpl) FindById(userId string) (*entities.User, error) {
	user, err := h.repositories.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *UserServiceImpl) FindList() *entities.UserList {
	users := h.repositories.GetListUser()
	return users
}
