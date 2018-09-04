package db

import (
	"database/sql"
)

// Wrapper -
type Wrapper struct {
	*sql.DB
	Connection string
}

func NewWrapper(connection string) *Wrapper {
	return &Wrapper{
		Connection: connection,
	}
}

// Init cache instance
func (o *Wrapper) Init() (*Wrapper, error) {
	var err error

	o.DB, err = sql.Open("sqlite3", o.Connection)

	return o, err
}

func (o *Wrapper) Ping() error {
	if err := o.DB.Ping(); err != nil {
		return err
	}

	return nil
}
