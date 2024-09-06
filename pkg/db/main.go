package db

import (
	"aura-test/pkg/config"
	"aura-test/pkg/log"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CreateConnection the instance of database via sqlx
func CreateConnection() *sqlx.DB {
	connValue := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=UTC&parseTime=True",
		config.EnvForge().GetString("DATABASE.CONNECT.ACCOUNT"),
		config.EnvForge().GetString("DATABASE.CONNECT.PASSWORD"),
		config.EnvForge().GetString("DATABASE.CONNECT.HOST"),
		config.EnvForge().GetString("DATABASE.CONNECT.PORT"),
		config.EnvForge().GetString("DATABASE.CONNECT.DB"),
	)

	db, err := sqlx.Open("mysql", connValue)
	if err != nil {
		log.Error("Connection to database failed. err ", err, " , connection value: ", connValue)
		return nil
	}
	db.SetMaxOpenConns(config.ConForge().GetInt("db.max_open_conns"))
	db.SetMaxIdleConns(config.ConForge().GetInt("db.max_idle_conns"))
	db.SetConnMaxLifetime(config.ConForge().GetDuration("db.conn_max_life_time") * time.Second)

	return db
}
