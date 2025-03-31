package auth_handler

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-google-openid/internal/core/domain/auth"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/http/dto"
)

type JwtHandler struct {
	Svc auth.JwtService
}

func NewJwtHandler(s auth.JwtService) *JwtHandler {
	return &JwtHandler{s}
}

func (h *JwtHandler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := h.Svc.Login(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(&dto.FailedResponse{
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&dto.Response{
		Data: dto.LoginResponse{
			Token:     token,
			TokenType: "Bearer",
		},
	})

}

func (h *JwtHandler) Register(w http.ResponseWriter, r *http.Request) {
	token, err := h.Svc.Register(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(&dto.FailedResponse{
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&dto.Response{
		Data: dto.RegisterResponse{
			Token:     token,
			TokenType: "Bearer",
		},
	})
}
