package app

import (
	"context"
	"log"

	"github.com/kenyako/chat-server/internal/client/db"
	"github.com/kenyako/chat-server/internal/client/db/pg"
	"github.com/kenyako/chat-server/internal/client/db/transaction"
	"github.com/kenyako/chat-server/internal/closer"
	"github.com/kenyako/chat-server/internal/config"
	"github.com/kenyako/chat-server/internal/config/env"
	"github.com/kenyako/chat-server/internal/repository"
	"github.com/kenyako/chat-server/internal/service"

	chatAPI "github.com/kenyako/chat-server/internal/api/chat"
	chatRepo "github.com/kenyako/chat-server/internal/repository/chat"
	chatServ "github.com/kenyako/chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	chatRepository repository.ChatAPIRepo
	chatService    service.ChatAPIService

	chatImpl *chatAPI.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {

	if s.pgConfig == nil {

		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {

	if s.grpcConfig == nil {

		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get gRPC config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBCClient(ctx context.Context) db.Client {

	if s.dbClient == nil {

		client, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create dbc client: %v", err)
		}

		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping db failed: %v", err)
		}
		closer.Add(client.Close)

		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {

	if s.txManager == nil {

		s.txManager = transaction.NewTransactionManager(s.DBCClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatAPIRepo {

	if s.chatRepository == nil {

		s.chatRepository = chatRepo.NewRepository(s.DBCClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatAPIService {

	if s.chatService == nil {

		s.chatService = chatServ.NewServiceChat(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chatAPI.Implementation {

	if s.chatImpl == nil {

		s.chatImpl = chatAPI.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
