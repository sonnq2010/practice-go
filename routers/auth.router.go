package routers

import (
	"net/http"
	"practice-go/controllers"
	"practice-go/repos"
	"practice-go/services"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(c controllers.IAuthController) http.Handler {
	r := chi.NewRouter()

	r.Get("/signup", c.SignUp)
	r.Get("/signin", c.SignIn)

	return r
}

func NewAuthRouter() http.Handler {
	return AuthRouter(controllers.NewAuthController(services.NewAuthService(repos.NewUserRepo())))
}
