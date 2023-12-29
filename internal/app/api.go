package app

import (
	"log/slog"
	"net/http"
	"os"

	"rest-service/internal/rest/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	router *chi.Mux
}

func New() *App {
	slog.Info("Starting app...", "user", os.Getenv("USER"))

	return &App{
		router: chi.NewRouter(),
	}
}

func (a *App) RunHTTPServer() {
	a.registerRoutes()
	http.ListenAndServe(":3333", a.router)

}

func (a *App) registerRoutes() {
	// Add middlewares
	a.router.Use(middleware.RequestID)
	//a.router.Use(middleware.Timeout(15 * time.Second))

	a.router.Get("/hello", handlers.HelloWorldHandler)
	a.router.Get("/random", handlers.RandomHandler)

}
