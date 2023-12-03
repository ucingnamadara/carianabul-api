package dto

import "dana/anabul-rest-api/src/entities"

type LoginFormRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	User  *entities.User `json:"user"`
	Token Token          `json:"token"`
}
