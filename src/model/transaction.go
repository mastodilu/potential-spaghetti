package model

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mastodilu/obsidian-finances/model/transaction"
)

type Record struct {
	TransactionType transaction.TransactionType
	amount          float32
	MoneyFrom       string
	MoneyTo         string
	When            time.Time
	Categories      []string
	Person          string
	Description     string
}

const (
	TR_TYPE     = "transaction_type::"
	AMOUNT      = "amount::"
	MONEYFROM   = "money-from::"
	MONEYTO     = "money-to::"
	WHEN        = "when::"
	CATEGORIES  = "categories::"
	PERSON      = "person::"
	DESCRIPTION = "description::"
)

const (
	DATE_LAYOUT = "2006/01/02"
)

func RecordFromFile(path string) (Record, error) {
	md := Record{}

	bb, err := os.ReadFile(path)
	if err != nil {
		return md, fmt.Errorf("error reading markdown file: %v", err)
	}

	lines := strings.Split(string(bb), "\n")

	for _, l := range lines {
		// exit if the previous iteration generated an error
		if err != nil {
			return md, err
		}

		switch {
		// go to next line if string is not a header
		case !strings.HasPrefix(l, "#"):
			continue
		case strings.Contains(l, TR_TYPE):
			md.TransactionType, err = transaction.GetTransactionType(
				getStringAfterPrefix(l, TR_TYPE),
			)
		case strings.Contains(l, AMOUNT):
			md.amount, err = _getAmount(getStringAfterPrefix(l, AMOUNT))
		case strings.Contains(l, MONEYFROM):
			md.MoneyFrom = _getMoneyFrom(getStringAfterPrefix(l, MONEYFROM))
		case strings.Contains(l, MONEYTO):
			md.MoneyTo = _getMoneyTo(getStringAfterPrefix(l, MONEYTO))
		case strings.Contains(l, WHEN):
			md.When, err = _getWhen(getStringAfterPrefix(l, WHEN))
		case strings.Contains(l, CATEGORIES):
			md.Categories = _getCategories(getStringAfterPrefix(l, CATEGORIES))
		case strings.Contains(l, PERSON):
			md.Person = _getPerson(getStringAfterPrefix(l, PERSON))
		case strings.Contains(l, DESCRIPTION):
			md.Description = _getDescription(getStringAfterPrefix(l, DESCRIPTION))
		}
	}

	return md, nil
}

func _getAmount(s string) (float32, error) {
	n, err := strconv.ParseFloat(s, 32)
	if err != nil {
		err = fmt.Errorf("error getting the amount from string '%s': %w", s, err)
	}
	return float32(n), err
}
func _getMoneyFrom(s string) string {
	return s
}
func _getMoneyTo(s string) string {
	return s
}
func _getWhen(s string) (time.Time, error) {
	s = strings.ReplaceAll(s, "#", "")
	return time.Parse(DATE_LAYOUT, s)
}
func _getCategories(s string) []string {
	return strings.Split(s, " ")
}
func _getPerson(s string) string {
	return s
}
func _getDescription(s string) string {
	return s
}

func getStringAfterPrefix(s, prefix string) string {
	begin := strings.Index(s, prefix)
	if begin == -1 {
		return ""
	}

	return strings.TrimSpace(s[begin+len(prefix):])
}

func (r *Record) Amount() float32 {
	switch r.TransactionType {
	case transaction.OUT:
		return -r.amount
	default:
		return r.amount
	}
}

func (r *Record) String() string {
	return fmt.Sprintf("TransactionType %v\nAmount %.2f\nMoneyFrom %v\nMoneyTo %v\nWhen %v\nCategories %v\nPerson %v\nDescription %v\n",
		r.TransactionType,
		r.Amount(),
		r.MoneyFrom,
		r.MoneyTo,
		r.When.Local().Format(DATE_LAYOUT),
		r.Categories,
		r.Person,
		r.Description,
	)
}
