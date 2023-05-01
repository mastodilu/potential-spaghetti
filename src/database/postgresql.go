package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	db *pgx.Conn
	once sync.Once
}

var (
	instance DB
)

func _establishConnection() {
	log.Println("connecting to postgresql")

	var err error
	url := os.Getenv("DATABASE_URL")

	instance.db, err = pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalln("Unable to connect to database.", err)
	}
	// defer instance.db.Close(context.Background())
}

func GetInstance() *DB {
	instance.once.Do(_establishConnection)
	return &instance
}

func (d *DB)Close() {
	if err := d.db.Close(context.Background()); err != nil {
		log.Fatalln(err)
	}
	instance = DB{}
}