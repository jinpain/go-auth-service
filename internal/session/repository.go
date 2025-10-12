package session

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinpain/go-auth-service/pkg/database/postgres"
	"github.com/jinpain/go-auth-service/pkg/sqlstore"
)

type Repository struct {
	pgxAdapter *postgres.PgxAdapter
	sqlStore   *sqlstore.SqlStore
}

func NewRepository(pgxAdapter *postgres.PgxAdapter, sqlStore *sqlstore.SqlStore) *Repository {
	return &Repository{
		pgxAdapter: pgxAdapter,
		sqlStore:   sqlStore,
	}
}

func (r *Repository) CreateSession(ctx context.Context, session *Model) error {
	query, err := r.sqlStore.GetQuery("create_session.sql")
	if err != nil {
		return err
	}

	err = r.pgxAdapter.TxOrDb(ctx).QueryRow(ctx, query,
		session.UserID,
		session.Device,
		session.IpAddress,
	).Scan(&session.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) RevokeSession(ctx context.Context, userID, sessionID uuid.UUID) error {
	query, err := r.sqlStore.GetQuery("revoke_session.sql")
	if err != nil {
		return err
	}

	tag, err := r.pgxAdapter.TxOrDb(ctx).Exec(ctx, query, userID, sessionID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("invalid session")
	}

	return nil
}
