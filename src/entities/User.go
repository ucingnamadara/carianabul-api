package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID          string    `json:"id" gorm:"primaryKey default:uuid_generate_v4() not null"`
	FullName    string    `json:"fullName" gorm:"not null"`
	PhoneNumber string    `json:"phoneNumber" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null"`
	Password    string    `json:"password" gorm:"not null"`
	Salt		string    `json:"salt" gorm:"not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type UserList []*User

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
