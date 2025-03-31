package auth_handler

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-google-openid/internal/core/domain/auth"
	"github.com/ilyasa1211/go-google-openid/internal/infrastructure/http/dto"
)

type OpenIDHandler struct {
	Svc auth.OpenIdService
}

func NewOpenIDHandler(s auth.OpenIdService) *OpenIDHandler {
	return &OpenIDHandler{s}
}

func (h *OpenIDHandler) Login(w http.ResponseWriter, r *http.Request) {
	url, err := h.Svc.GetLoginUrl()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(&dto.FailedResponse{
			Message: err.Error(),
		})

		return
	}

	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

func (h *OpenIDHandler) HandleLoginCallback(w http.ResponseWriter, r *http.Request) {
	token, err := h.Svc.HandleLoginCallback(r)

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
