package user

import (
	"context"

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

func (r *Repository) CreateUser(ctx context.Context, user *Model) error {
	query, err := r.sqlStore.GetQuery("create_user.sql")
	if err != nil {
		return err
	}
	err = r.pgxAdapter.TxOrDb(ctx).QueryRow(ctx, query,
		user.Email,
		user.Phone,
		user.Password,
	).Scan(&user.ID)

	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*Model, error) {
	query, err := r.sqlStore.GetQuery("get_user_by_email.sql")
	if err != nil {
		return nil, err
	}

	var user Model
	err = r.pgxAdapter.TxOrDb(ctx).QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Email,
		&user.Phone,
		&user.Password,
		&user.Verified,
		&user.Blocked,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) ExistsUserByEmail(ctx context.Context, email string) (bool, error) {
	query, err := r.sqlStore.GetQuery("exists_user_by_email.sql")
	if err != nil {
		return false, err
	}

	var exists bool
	err = r.pgxAdapter.TxOrDb(ctx).QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) SetUserVerified(ctx context.Context, userID uuid.UUID) error {
	query, err := r.sqlStore.GetQuery("set_user_verified.sql")
	if err != nil {
		return err
	}
	_, err = r.pgxAdapter.TxOrDb(ctx).Exec(ctx, query, userID)
	return err
}
