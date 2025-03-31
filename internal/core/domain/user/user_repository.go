package user

import (
	"github.com/ilyasa1211/go-google-openid/internal/application/dto"
)

type UserRepository interface {
	FindAll() []*User
	FindById(id string) *User
	FindByEmail(email string) *User
	Create(u *dto.CreateUserRequest) error
	UpdateById(id string, u *dto.UpdateUserRequest) error
	DeleteById(id string) error
}
