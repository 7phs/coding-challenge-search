package memory

import (
	"errors"

	"github.com/7phs/coding-challenge-search/model"
	log "github.com/sirupsen/logrus"
)

type ItemsSource interface {
	Load() (model.ItemsList, error)
}

type Items struct {
	source ItemsSource

	list model.ItemsList
}

func NewItems(source ItemsSource) *Items {
	return &Items{
		source: source,
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
	if err := o.reindex(); err != nil {
		return errors.New(logPrefix + " failed to reindex items, " + err.Error())
	}

	return nil
}

func (o *Items) load() (err error) {
	o.list, err = o.source.Load()

	return
}

func (o *Items) reindex() error {
	return nil
}

func (o *Items) List(filter *model.SearchFilter) (model.ItemsList, error) {
	return nil, nil
}
