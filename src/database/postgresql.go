package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DB struct {
	db *pgx.Conn
	once sync.Once
}

var (
	instance DB
)

const (
	UniqueViolation = "23505"
)

func _establishConnection() {
	log.Println("connecting to postgresql")

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("cannot parse database port: %v\n", err)
	}
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_USER"),
		os.Getenv("DB_USER_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	
	instance.db, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalln("Unable to connect to database.", err)
	}
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

// Exec calls Exec on the underlying database
func (d *DB) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error){
	log.Println(query)
	return d.db.Exec(ctx, query, args...)
}

// QueryRow calls QueryRow on the underlying database
func (d *DB)QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	log.Println(query)
	return d.db.QueryRow(ctx, query, args...)
}

// Query calls Query on the underlying database
func (d *DB)Query(ctx context.Context, query string, args ...any) (pgx.Rows, error){
	log.Println(query)
	return d.db.Query(ctx,query, args...)
}

