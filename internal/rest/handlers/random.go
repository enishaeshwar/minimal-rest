package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"rest-service/internal/service/random"
)

type RandomHandler struct {
	Service *random.Service
}

type Response struct {
	Values []int `json:"values"`
}

func (h *RandomHandler) RandomHandle(w http.ResponseWriter, r *http.Request) {

	n := r.URL.Query().Get("value")

	val, err := strconv.Atoi(n)
	if err != nil {
		slog.Error("Error in handling request", "n", n, "val", val)
		w.Write([]byte("error"))
		return
	}

	res := h.Service.RandomNumbers(val)

	//resS := strings.Trim(strings.Join(strings.Split(fmt.Sprint(res), " "), ","), "[]")

	j, _ := json.Marshal(Response{Values: res})

	if _, err := w.Write(j); err != nil {
		slog.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
