package db

import (
	"github.com/7phs/coding-challenge-search/db/memory"
	"github.com/7phs/coding-challenge-search/db/sqlite"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	MemoryItems *memory.Items
)

func Init(connection string) {
	log.Info("DB: init a db")
	if err := sqlite.Init(connection); err != nil {
		log.Fatal("DB: init a db, failed: ", err)
	}

	MemoryItems = memory.NewItems(sqlite.ItemsSource)

	log.Info("DB: memory, init DB")
	if err := MemoryItems.Init(); err != nil {
		log.Fatal("DB: memory, init DB - failed: ", err)
	}
}

func Shutdown() {
	log.Info("DB: shutdown")
	sqlite.Shutdown()
}
