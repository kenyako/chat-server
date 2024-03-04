package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	desc "github.com/kenyako/chat-server/pkg/chat_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=5433 dbname=chat-service user=chat-user password=chat-password sslmode=disable"
)

type server struct {
	desc.UnimplementedChatAPIServer

	db *pgxpool.Pool
	qb sq.StatementBuilderType
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	chatTitle := req.GetTitle()
	members := req.GetUsernames()

	builderInsert := s.qb.Insert("chats").
		Columns("title").
		Values(chatTitle).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to get rows: %v", err)
	}
	defer rows.Close()

	chatID, err := pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		log.Fatalf("failed to get chatID: %v", err)
	}

	builderInsert = s.qb.Insert("users_chats").
		Columns("chat_id", "user_id")

	// генерация ID пользователей, позже будет добавлена ручка для получения ID из базы auth
	for _, username := range members {
		userID := int64(gofakeit.Uint8()) + int64(len(username))

		log.Printf("user with id: %d added in chat", userID)

		builderInsert = builderInsert.Values(chatID, userID)
	}

	query, args, err = builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to added data in db: %v", err)
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	chatID := req.GetId()

	builderDelete := s.qb.Delete("users_chats").
		Where(sq.Eq{"chat_id": chatID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to biuld delete query UsersChats: %v", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete UsersChats: %v", err)
	}

	builderDelete = s.qb.Delete("chat_messages").
		Where(sq.Eq{"chat_id": chatID})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to biuld delete query Messages: %v", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete Messages: %v", err)
	}

	builderDelete = s.qb.Delete("chats").
		Where(sq.Eq{"id": chatID})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to biuld delete query Chats: %v", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete Chats: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendRequest) (*emptypb.Empty, error) {

	chatID := req.GetChatID()
	from := req.GetFrom()
	text := req.GetText()

	builderInsert := s.qb.Insert("chat_messages").
		Columns("chat_id", "user_id", "text", "time_sent").
		Values(chatID, from, text, time.Now())

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build insert Messages: %v", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to insert Messages: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatAPIServer(s, &server{
		db: pool,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
