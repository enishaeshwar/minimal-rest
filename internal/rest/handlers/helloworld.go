package handlers

import (
	"log/slog"
	"net/http"

	"rest-service/internal/service/helloworld"
)

type HelloWorldHandler struct {
	Service *helloworld.Service
}

func (h *HelloWorldHandler) HelloWorldHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if ctx.Err() == nil {
		slog.Info("Performing some large calculation")
		w.Write([]byte(h.Service.HelloWorld()))
	} else {
		slog.Info("Context cancelled. Not performing large calculation.")
		return
	}
}
