package main

import (
	"rest-service/internal/app"
	"rest-service/internal/logger"
)

func main() {
	logger.Init()
	a := app.New()
	a.RunHTTPServer()
}
