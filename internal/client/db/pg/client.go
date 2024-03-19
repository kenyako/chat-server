package pg

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kenyako/chat-server/internal/client/db"
)

type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, errors.New("failed to connect to db")
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (pc *pgClient) DB() db.DB {
	return pc.masterDBC
}

func (pc *pgClient) Close() error {
	if pc.masterDBC != nil {
		pc.masterDBC.Close()
	}

	return nil
}
