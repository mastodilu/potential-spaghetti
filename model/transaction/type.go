package transaction

import (
	"fmt"
	"strings"
)

type TransactionType uint8

const (
	NONE TransactionType = iota
	OUT
	CREDIT
	DEBIT
	TRANSFER
	WITHDRAW
	IN
	IN_REFUND
	IN_INCOMING
	IN_MISSING
)

func GetTransactionType(s string) (TransactionType, error) {
	s = strings.ToLower(s)
	var err error

	switch {
	case strings.Contains(s, "spesa/type/out"):
		return OUT, err
	case strings.Contains(s, "spesa/type/credit"):
		return CREDIT, err
	case strings.Contains(s, "spesa/type/transfer"):
		return TRANSFER, err
	case strings.Contains(s, "spesa/type/debit"):
		return DEBIT, err
	case strings.Contains(s, "spesa/type/withdraw"):
		return WITHDRAW, err

	case strings.Contains(s, "spesa/type/in"):
		substring := s[len("#spesa/type/in"):]
		switch {
		case strings.Contains(substring, "refund"):
			return IN_REFUND, err
		case strings.Contains(substring, "incoming"):
			return IN_INCOMING, err
		case strings.Contains(substring, "missing"):
			return IN_MISSING, err
		}
	}
	return NONE, fmt.Errorf("transaction_type not found in string '%s'", s)
}
