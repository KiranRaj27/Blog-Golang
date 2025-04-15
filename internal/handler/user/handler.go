// internal/handler/user/handler.go
package user

import (
	"encoding/json"
	"net/http"

	domain "github.com/kiranraj27/blog-golang/internal/domain/user"
	"github.com/kiranraj27/blog-golang/internal/usecase/user"
	"github.com/kiranraj27/blog-golang/pkg/jwt"
)

type Handler struct {
	svc *user.Service
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc}
}

type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	u := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.svc.Register(u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered",
		"user":    u,
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.svc.Authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Generate(user.Email)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
