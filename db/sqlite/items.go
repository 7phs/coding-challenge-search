package sqlite

import (
	"database/sql"
	"errors"

	"github.com/7phs/coding-challenge-search/model"
)

type Items struct {
	db *sql.DB
}

func NewItems(db *sql.DB) *Items {
	return &Items{
		db: db,
	}
}

func (o *Items) Load() (model.ItemsList, error) {
	var (
		load      = QueryItemsLoad{}
		logPrefix = "sqlite/items: load"
	)

	rows, err := o.db.Query(load.Query(), load.Bind()...)
	if err != nil {
		return nil, errors.New(logPrefix + " - failed to execute a query, " + err.Error())
	}
	defer rows.Close()

	result := make(model.ItemsList, 0, 100)

	for rows.Next() {
		record := model.Item{}

		if err := load.Scan(rows, &record); err != nil {
			return nil, errors.New(logPrefix + " - failed to scan a row, " + err.Error())
		}

		result.Add(&record)
	}

	return result, nil
}
