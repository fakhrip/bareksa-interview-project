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

	// Create all tables if not exist
	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*domain.News)(nil)).IfNotExists().WithForeignKeys().Exec(ctx)
	if err != nil {
		panic(err)
	}

	_, err = db.NewCreateTable().Model((*domain.Topics)(nil)).IfNotExists().WithForeignKeys().Exec(ctx)
	if err != nil {
		panic(err)
	}

	return db
}
