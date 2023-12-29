package handlers

import (
	"log/slog"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if ctx.Err() == nil {
		slog.Info("Performing some large calculation")
		w.Write([]byte("hello world"))
	} else {
		slog.Info("Context cancelled. Not performing large calculation.")
		return
	}
}
