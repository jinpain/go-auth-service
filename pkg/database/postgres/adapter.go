package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Querier interface {
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
}

type PgxAdapter struct {
	pool *pgxpool.Pool
}

func NewPgxAdapter(connStr string) (*PgxAdapter, error) {
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &PgxAdapter{pool: pool}, nil

}

func (pg *PgxAdapter) TxOrDb(ctx context.Context) Querier {
	tx, ok := ctx.Value(transactionKey).(pgx.Tx)
	if !ok {
		return pg.pool
	}

	return tx
}

func (p *PgxAdapter) Close() {
	if p.pool != nil {
		p.pool.Close()
	}
}
