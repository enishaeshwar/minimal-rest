package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	eni1 := make(chan int)

	go dummyGoRoutine("Eni1", eni1)
	go dummyGoRoutine("Eni2", eni1)

	<-eni1
	<-eni1

	if ctx.Err() == nil {
		slog.Info("Performing some large calculation")
		w.Write([]byte("hello world"))
	} else {
		slog.Info("Context cancelled. Not performing large calculation.")
		return
	}
}

func dummyGoRoutine(name string, c chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(name, ":", i)
	}
	c <- 1000
}
