package model

import (
	"../../config"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const (
	tableNameSignalEvents = "signal_events"
)

var DbConnection *sql.DB

// create db schema
func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// create table
	cmd := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            time DATETIME PRIMARY KEY NOT NULL,
            product_code STRING,
            side STRING,
            price FLOAT,
            size FLOAT)`, tableNameSignalEvents)
	DbConnection.Exec(cmd)

	// create table
	for _, duration := range config.Config.Durations {
		tableName := GetCandleTableName(config.Config.ProductCode, duration)
		c := fmt.Sprintf(`
            CREATE TABLE IF NOT EXISTS %s (
            time DATETIME PRIMARY KEY NOT NULL,
            open FLOAT,
            close FLOAT,
            high FLOAT,
            low FLOAT,
			volume FLOAT)`, tableName)
		DbConnection.Exec(c)
	}
}

func GetCandleTableName(productCode string, duration time.Duration) string {
	return fmt.Sprintf("%s_%s", productCode, duration)
}
