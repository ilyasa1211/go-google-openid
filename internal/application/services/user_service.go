package services

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-google-openid/internal/application/dto"
	"github.com/ilyasa1211/go-google-openid/internal/core/domain/user"
	"github.com/ilyasa1211/go-google-openid/internal/utils"
)

type UserService struct {
	r user.UserRepository
}

func NewUserService(r user.UserRepository) *UserService {
	return &UserService{r}
}

func (s *UserService) FindAll() []*user.User {
	return s.r.FindAll()
}
func (s *UserService) FindById(r *http.Request) *user.User {
	id := r.PathValue("id")

	return s.r.FindById(id)
}
func (s *UserService) Create(r *http.Request) error {
	var dto dto.CreateUserRequest

	json.NewDecoder(r.Body).Decode(&dto)

	dto.Password = utils.HashPass(dto.Password)

	return s.r.Create(&dto)
}
func (s *UserService) UpdateById(r *http.Request) error {
	id := r.PathValue("id")

	var dto dto.UpdateUserRequest
	json.NewDecoder(r.Body).Decode(&dto)

	return s.r.UpdateById(id, &dto)
}
func (s *UserService) DeleteById(r *http.Request) error {
	id := r.PathValue("id")

	return s.r.DeleteById(id)
}
