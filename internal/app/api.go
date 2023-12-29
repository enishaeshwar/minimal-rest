package app

import (
	"log/slog"
	"net/http"
	"os"

	"rest-service/internal/rest/handlers"
	"rest-service/internal/service/helloworld"
	"rest-service/internal/service/random"

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
	hh := handlers.HelloWorldHandler{HelloSrv: helloworld.New()}
	rh := handlers.RandomHandler{RandSrv: random.New()}

	a.registerRoutes(hh, rh)

	http.ListenAndServe(":3333", a.router)

}

func (a *App) registerRoutes(hh handlers.HelloWorldHandler, rh handlers.RandomHandler) {
	// Add middlewares
	a.router.Use(middleware.RequestID)
	//a.router.Use(middleware.Timeout(15 * time.Second))

	a.router.Get("/hello", hh.HelloWorldHandle)
	a.router.Get("/random", rh.RandomHandle)

}
