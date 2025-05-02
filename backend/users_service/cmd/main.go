package main

import (
	"context"
	"log"

	"users_service/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("ошибка инициализации сервиса: %s", err.Error())
	}

	err = a.RunGRPCServer()
	if err != nil {
		log.Fatalf("ошибка запуска сервиса: %s", err.Error())
	}
}
