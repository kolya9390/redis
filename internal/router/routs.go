package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	reverproxy "studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/tools/reverProxy"
	swaggerui "studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/tools/swaggerUI"
	"studentgit.kata.academy/Nikolai/historysearch/internal/infrastructure/tools/token"
	"studentgit.kata.academy/Nikolai/historysearch/internal/modules"
)

func NewApiRouter(controllers modules.Controller) http.Handler {
	r := chi.NewRouter()


	proxy := reverproxy.NewReverseProxy("hugo", "1313")


	r.Use(middleware.Logger)
	r.Use(proxy.ReverseProxy)

	//SwaggerUI
	r.Get("/swagger", swaggerui.SwaggerUI)

	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./app/public"))).ServeHTTP(w, r)
	})

	// API
	r.Route("/api", func(r chi.Router) {
		authController := controllers.AuthController
		geo := controllers.GeoController

		r.Post("/address/search",geo.SearchAPI)
		r.Post("/address/geocode",geo.GeocodeAPI)

		// Group users

		
		r.Post("/login", authController.Login)

		r.Post("/register", authController.Register)

		// Group Adress
		r.Route("/address", func(r chi.Router) {
			token := token.TokenAuthorization

			r.Use(jwtauth.Verifier(token))

			r.Use(jwtauth.Authenticator(token))

		})
	})

	return r
}