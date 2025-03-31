package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-google-openid/internal/application/services"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/http/dto"
)

type UserHandler struct {
	Svc *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.FindAll()

	resp := &dto.Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.FindById(r)

	resp := &dto.Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.Create(r)

	resp := &dto.Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.UpdateById(r)

	resp := &dto.Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.DeleteById(r)

	resp := &dto.Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
