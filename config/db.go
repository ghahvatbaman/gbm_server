package config

import (
	"database/sql"
	"errors"

	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func SetupDb() (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, viper.Get("db").(string))
	if err != nil {
		return nil, errors.New("database not initialized")
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())

	return db, nil
}
