package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	userHandler "github.com/kiranraj27/blog-golang/internal/handler/user"
	userRepo "github.com/kiranraj27/blog-golang/internal/repository/user"
	userUsecase "github.com/kiranraj27/blog-golang/internal/usecase/user"
)

func New() http.Handler {
	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	repo := userRepo.NewUserRepo()
	svc := userUsecase.NewService(repo)
	handler := userHandler.NewHandler(svc)

	r.Post("/login", handler.Login)
	r.Post("/register", handler.Register)

	return r
}
