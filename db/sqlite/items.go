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
	index := int64(1)

	for rows.Next() {
		record := model.Item{
			Id: index,
		}

		if err := load.Scan(rows, &record); err != nil {
			return nil, errors.New(logPrefix + " - failed to scan a row, " + err.Error())
		}

		record.Location.PreCalc()

		index++
		result.Add(&record)
	}

	return result, nil
}
