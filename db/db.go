package db

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var DB *Wrapper

func Init(connection string) {
	var (
		err error
	)

	log.Info("DB: create a connection pool")
	DB, err = NewWrapper(connection).
		Init()
	if err != nil {
		log.Fatal("DB: failed to create a connection pool: ", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("DB: check DB connection - failed: ", err)
	} else {
		log.Info("DB: check DB connection - successful")
	}
}

func Shutdown() {
}
