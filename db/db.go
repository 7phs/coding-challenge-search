package db

import (
	"github.com/7phs/coding-challenge-search/db/cache"
	"github.com/7phs/coding-challenge-search/db/index"
	"github.com/7phs/coding-challenge-search/db/memory"
	"github.com/7phs/coding-challenge-search/db/sqlite"
	"github.com/7phs/coding-challenge-search/model"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var (
	MemoryItems *cache.Items
)

type Dependencies struct {
	Lem model.Lemmer
}

func Init(connection string, dep Dependencies) {
	log.Info("DB: init a db")
	if err := sqlite.Init(connection); err != nil {
		log.Fatal("DB: init a db, failed: ", err)
	}

	log.Info("DB: memory, init DB")
	items := memory.NewItems(sqlite.ItemsSource,
		index.NewIndexWords(dep.Lem),
		index.NewIndexTiles())
	if err := items.Init(); err != nil {
		log.Fatal("DB: memory, init DB - failed: ", err)
	}

	log.Info("DB: memory, init a DB cache layer")
	MemoryItems = cache.NewItems(items)
}

func Shutdown() {
	log.Info("DB: shutdown")
	sqlite.Shutdown()
}
