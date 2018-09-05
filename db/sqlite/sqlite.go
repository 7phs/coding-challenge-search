package sqlite

import (
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB

	ItemsSource *Items
)

func Init(connection string) (err error) {
	log.Info("DB: sqlite - open a db '" + connection + "'")
	db, err = sql.Open("sqlite3", connection)
	if err != nil {
		return errors.New("sqlite: failed to open a db, " + err.Error())
	}

	log.Info("DB: sqlite - ping a connection")
	err = db.Ping()
	if err != nil {
		return errors.New("sqlite: failed to ping a connection, " + err.Error())
	}

	log.Info("DB: sqlite - create an items data source")
	ItemsSource = NewItems(db)

	return nil
}

func Shutdown() {
	log.Info("DB: sqlite - shutdown")
	if db != nil {
		db.Close()
	}
}
