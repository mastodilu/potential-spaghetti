package transactionhandler

import (
	"context"

	"github.com/mastodilu/obsidian-finances/database"
	model "github.com/mastodilu/obsidian-finances/types"
)

// type transaction is an alias over model.Transaction
type transaction model.Transaction

// var db SQLDBInstance

// const (
// 	custom errors...
// )

// import transaction.Transaction

func (t *transaction) InsertOne() error {
	panic("method not implemented yet! 😱")
	db := database.GetInstance()
	query := "INSERT ... INTO ... "

	_, err := db.Exec(context.Background(), query)
	return err
}