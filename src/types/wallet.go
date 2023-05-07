package types

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mastodilu/obsidian-finances/database"
)

const (
	PREFIX_SPESA_FROM = "#spesa/from/"
	PREFIX_SPESA_TO = "#spesa/to/"
)

type Wallet string

var (
	wallets = []Wallet{}
	rx = regexp.MustCompile(`#spesa\/(from|to)\/?`)
)

func PrintWallets() {
	log.Println("🐒")
	log.Println(wallets)
}

// walletFromString cleans the label and
// returns a Wallet from it
func walletFromString(label string) Wallet{
	shortLabel := strings.TrimSpace(rx.ReplaceAllString(label, ""))
	w := Wallet(shortLabel)
	return w
}

func (w *Wallet) isInvalid() bool {
	return strings.Compare(string(*w), "") == 0
}

// AddWallet add a cleaned wallet to the DB.
// If successful, it also adds it in the local list of wallets.
func AddWallet(wallet Wallet) {
	if wallet.In(wallets) || wallet.isInvalid() {
		return
	}
	
	db := database.GetInstance()
	query := fmt.Sprintf("INSERT INTO schema.wallet (label) VALUES ('%s');", wallet)
	_, err := db.Exec(context.Background(), query)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && strings.Compare(pgErr.Code, database.UniqueViolation) == 0 {
			// nothing to do, wallet already in database
		} else {
			log.Fatalf("error adding wallet to DB: %v", err)
		}
	}
	wallets = append(wallets, wallet)
}

func (w Wallet)In(ww []Wallet) bool {
	for _, wallet := range ww {
		if strings.Compare(string(wallet), string(w)) == 0 {
			return true
		}
	}
	return false
}