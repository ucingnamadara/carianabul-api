package utils

import (
	"dana/anabul-rest-api/src/dto"
	"dana/anabul-rest-api/src/entities"
	"os"

	"time"

	"github.com/golang-jwt/jwt"
)

type JwtUtils interface {
	JwtSign(user *entities.User) (dto.Token, error)
}

type JwtUtilsImpl struct {
}

func NewJwtUtils() *JwtUtilsImpl {
	return &JwtUtilsImpl{}
}

var secretKay = []byte(os.Getenv("SECRET_KEY"))

func JwtSign(user *entities.User) (dto.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["phoneNumber"] = user.PhoneNumber
	claims["fullName"] = user.FullName
	claims["tokenType"] = "accessToken"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString(secretKay)

	if err != nil {
		return dto.Token{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["tokenType"] = "refreshToken"
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshTokenString, err := refreshToken.SignedString(secretKay)

	if err != nil {
		return dto.Token{}, err
	}
	return dto.Token{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
