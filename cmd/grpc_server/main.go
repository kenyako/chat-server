package main

import (
	"context"
	"log"

	"github.com/kenyako/chat-server/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to create new App: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run gRPC server: %v", err)
	}
}
