package handlers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mastodilu/obsidian-finances/database"
	"github.com/mastodilu/obsidian-finances/types"
)

// WalletHandler is the only handlers capable of
// handling wallets in the database
type WalletHandler struct {}

// DBWallet alias for the wallet type
type DBWallet types.Wallet

// Insert inserts the list of wallets into the database
func (wh *WalletHandler) Insert(db database.DB, wallets ...types.Wallet) []error {
	var errors []error
	for _, w := range wallets {
		err := DBWallet(w).Insert(db)
		errors = append(errors, err)
	}
	return errors // 💚
}

// isInvalid returns true if the wallet is an empty string
func (dbw DBWallet) isInvalid() bool {
	return strings.Compare(string(dbw), "") == 0
}

func (dbw DBWallet)Insert(db database.DB) error {
	// do not insert a wallet with empty label
	if dbw.isInvalid() {
		return nil // 💚
	}
	
	query := fmt.Sprintf("INSERT INTO schema.wallet (label) VALUES ('%s');", dbw)

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && strings.Compare(pgErr.Code, database.UniqueViolation) == 0 {
			// wallet already in database
		} else {
			return fmt.Errorf("error adding wallet to DB: %v", err)
		}
	}
	return nil // 💚
}