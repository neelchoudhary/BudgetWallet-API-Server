package postgresql

import (
	"context"
	"database/sql"
)

// TxRepository interface
type TxRepository interface {
	StartTx(ctx context.Context) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
}

// NewTxRepository returns a new instance of a tx repository.
func NewTxRepository(db *sql.DB) TxRepository {
	return &txRepository{db: db}
}

type txRepository struct {
	db *sql.DB
}

func (r *txRepository) StartTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *txRepository) CommitTx(tx *sql.Tx) error {
	err := tx.Commit()
	return err
}
