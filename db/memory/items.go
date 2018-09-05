package memory

import (
	"errors"
	"os"
	"time"

	"github.com/7phs/coding-challenge-search/db/index"
	"github.com/7phs/coding-challenge-search/model"
	log "github.com/sirupsen/logrus"
)

type ItemsSource interface {
	Load() (model.ItemsList, error)
}

type Items struct {
	source ItemsSource

	list    model.ItemsList
	indexes []index.ItemIndex
}

func NewItems(source ItemsSource, indexes ...index.ItemIndex) *Items {
	return &Items{
		source:  source,
		indexes: indexes,
	}
}

func (o *Items) Init() error {
	logPrefix := "memory/items:"

	log.Info(logPrefix + " load items list from a data source")
	if err := o.load(); err != nil {
		return errors.New(logPrefix + " failed to load items list, " + err.Error())
	}
	log.Info(logPrefix+" loaded ", len(o.list), " records")

	log.Info(logPrefix + " reindexing records")
	start := time.Now()
	if err := o.reindex(); err != nil {
		return errors.New(logPrefix + " failed to reindex items, " + err.Error())
	}
	log.Info(logPrefix + " reindexing records for " + time.Since(start).String())

	return nil
}

func (o *Items) load() (err error) {
	o.list, err = o.source.Load()

	return
}

func (o *Items) reindex() error {
	for _, record := range o.list {
		for _, idx := range o.indexes {
			idx.Add(record)
		}
	}

	for _, idx := range o.indexes {
		idx.Finish()
	}

	return nil
}

func (o *Items) List(filter *model.SearchFilter, paging *model.Paging) (model.ItemsList, error) {
	var total index.Result

	for _, idx := range o.indexes {
		result, err := idx.Search(filter)
		if err != nil {
			if err != os.ErrInvalid && err != os.ErrNotExist {
				log.Error("memory/items: failed to request a search result, ", err)
			}

			continue
		}

		if total == nil {
			total = result
			total.SetMode(model.KeywordsModeAnd)
		} else {
			total = total.Reduce(result)
		}
	}

	if total == nil {
		return nil, errors.New("memory/items: not found")
	}

	result := total.Items(paging.Start, paging.Limit)

	if result == nil {
		return nil, errors.New("memory/items: start is out of bound")
	}

	return result, nil
}
