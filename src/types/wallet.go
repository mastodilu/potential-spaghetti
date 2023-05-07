package types

import (
	"log"
	"regexp"
	"strings"
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