package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"referalMS/internal/config"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewClient(ctx context.Context, pg config.StorageConfig) (pool *pgxpool.Pool, err error) {
	const op = "referalMS.pkg.client.postgres.NewClient"
	DSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", pg.User, pg.Password, pg.Host, pg.Port, pg.Database)

	err = ConnectWithRetry(func() error {
		ctx, cancel := context.WithTimeout(ctx, pg.RetryTimeout)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, DSN)
		if err != nil {
			return err
		}
		return nil
	}, pg.MaxRetry, pg.RetryTimeout)

	if err != nil {
		return pool, err
	}

	return pool, nil
}

func ConnectWithRetry(fn func() error, maxRetry int, timeout time.Duration) (err error) {
	for maxRetry > 0 {
		if err = fn(); err != nil {
			time.Sleep(timeout)
			maxRetry--
			continue
		}
		return nil
	}
	return
}
