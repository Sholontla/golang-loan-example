package api

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/projects/loans/utils/config"
)

func GetDbClient() *sqlx.DB {

	sslmode, postgresqlUsername, postgresqlPassword, postgresqlHost, postgresqlPort, postgresqlSchema := config.DBConfig()

	psqlConn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		postgresqlHost, postgresqlPort, postgresqlUsername, postgresqlPassword, postgresqlSchema, sslmode)

	Client, err := sqlx.Connect("pgx", psqlConn)
	if err != nil {
		panic(err)
	}

	Client.SetConnMaxLifetime(time.Minute * 3)
	Client.SetMaxOpenConns(10)
	Client.SetMaxIdleConns(10)

	return Client
}
