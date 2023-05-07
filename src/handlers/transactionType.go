package handlers

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mastodilu/obsidian-finances/database"
	"github.com/mastodilu/obsidian-finances/types"
)

// TransactionTypeHandler database handler for TransactionType
type TransactionTypeHandler struct{}

// DBTransactionType wrapper for TransactionType
type DBTransactionType struct {
	transactionType types.TransactionType
}

func (dbtt DBTransactionType) IsInvalid() bool {
	return dbtt.transactionType <= types.NONE || dbtt.transactionType > types.IN_MISSING
}

func (tth TransactionTypeHandler) Insert(db database.DB, types ...types.TransactionType) []error {
	var errors []error
	for _, tt := range types {
		dbtt := DBTransactionType{
			transactionType: tt,
		}
		err := dbtt.insertOne(db)
		errors = append(errors, err)
	}
	return errors
}

func (dbtt DBTransactionType) insertOne(db database.DB) error {
	if dbtt.IsInvalid() {
		return fmt.Errorf("invalid transaction type")
	}

	query := fmt.Sprintf("INSERT INTO schema.transaction_type(id, label) VALUES (%d, '%s');",
		dbtt.transactionType,
		dbtt.transactionType.String(),
	)

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// already existing in DB, nothing to do
		} else {
			return fmt.Errorf("error adding transactionType '%s' to the DB: %v", dbtt.transactionType.String(), err)
		}
	}
	return nil // 💚
}
