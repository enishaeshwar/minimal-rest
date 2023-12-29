package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"rest-service/internal/service/random"
)

type RandomHandler struct {
	RandSrv *random.Service
}

func (h *RandomHandler) RandomHandle(w http.ResponseWriter, r *http.Request) {

	n := r.URL.Query().Get("value")
	val, err := strconv.Atoi(n)
	if err != nil {
		slog.Error("Error in handling request", "val", val)
		w.Write([]byte("error"))
		return
	}

	res := h.RandSrv.RandomNumbers(val)

	resS := strings.Trim(
		strings.Join(
			strings.Split(fmt.Sprint(res), " "),
			","),
		"[]")

	w.Write([]byte(resS))
}
