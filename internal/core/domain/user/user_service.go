package user

import (
	"net/http"
)

type UserService interface {
	FindAll() []*User
	FindById(r *http.Request) *User
	Create(r *http.Request) error
	UpdateById(r *http.Request) error
	DeleteById(r *http.Request) error
}
