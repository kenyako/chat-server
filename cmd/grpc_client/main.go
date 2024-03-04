package main

import (
	"context"
	"log"
	"time"

	desc "github.com/kenyako/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

var usernames = []string{"Bob", "Alice", "Bruce", "Joshua", "Elen"}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connecrion: %v", err)
	}

	defer conn.Close()

	c := desc.NewChatAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Create(ctx, &desc.CreateRequest{Usernames: usernames, Title: "Best"})
	if err != nil {
		log.Fatalf("failed to create chat: %v", err)
	}

	log.Printf("create chat with ID: %d", r.GetId())
}
