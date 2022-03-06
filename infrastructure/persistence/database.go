package persistence

import (
	domain "bareksa-interview-project/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func CreateDatabase(dbPass string) *bun.DB {
	connectionString := fmt.Sprintf("postgres://postgres:%s@postgres:5432/postgres?sslmode=disable", dbPass)

	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())

	// Drop and create the table again (Refresh the whole database)
	err = db.ResetModel(context.Background(), (*domain.News)(nil), (*domain.Topics)(nil))
	if err != nil {
		panic(err)
	}

	return db
}
