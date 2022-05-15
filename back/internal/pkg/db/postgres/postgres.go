package database

import (
	"database/sql"
	"log"

	"github.com/glyphack/graphlq-golang/internal/config"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	config, err := config.LoadDBConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := sql.Open("postgres", config.Source)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
