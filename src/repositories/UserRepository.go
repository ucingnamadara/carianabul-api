package repositories

import (
	"dana/anabul-rest-api/src/dto"
	"dana/anabul-rest-api/src/entities"
	errors "dana/anabul-rest-api/src/templates/error"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetListUser() *entities.UserList
	GetUserById(userId string) (*entities.User, error)
	GetUserByEmail(username string) (*entities.User, error)
	Create(request dto.UserRegisterDTO) *entities.User
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repo *UserRepositoryImpl) GetListUser() *entities.UserList {
	result := entities.UserList{}
	repo.db.Find(&result)
	return &result
}

func (repo *UserRepositoryImpl) GetUserById(userId string) (*entities.User, error) {
	result := entities.User{}
	repo.db.Where("id = ?", userId).First(&result)
	if (entities.User{}) == result {
		return nil, errors.BadRequest("DATA_NOT_EXIST")
	}
	return &result, nil
}

func (repo *UserRepositoryImpl) GetUserByEmail(email string) (*entities.User, error) {
	result := entities.User{}
	repo.db.Where("email = ?", email).First(&result)
	if (entities.User{}) == result {
		return nil, errors.BadRequest("DATA_NOT_EXIST")
	}
	return &result, nil
}

func (repo *UserRepositoryImpl) Create(request dto.UserRegisterDTO) *entities.User {
	user := entities.User{
		FullName:    request.FullName,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	}
	repo.db.Create(&user)

	return &user
}
