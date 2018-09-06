package cache

import (
	"encoding/hex"
	"os"
	"sync"
	"time"

	"github.com/7phs/coding-challenge-search/model"
	log "github.com/sirupsen/logrus"
)

type ItemsSource interface {
	List(filter *model.SearchFilter, paging *model.Paging) (model.ItemsList, error)
}

type Items struct {
	sync.Map

	source ItemsSource
}

func NewItems(source ItemsSource) *Items {
	return &Items{
		source: source,
	}
}

func (o *Items) List(filter *model.SearchFilter, paging *model.Paging) (model.ItemsList, error) {
	key := hex.EncodeToString(filter.Hash(paging.Hash(nil)))

	logPrefix := "cache: " + filter.String()

	if data, ok := o.Load(key); ok {
		log.Info(logPrefix + " - found")

		switch v := data.(type) {
		case error:
			return nil, v
		case model.ItemsList:
			return v, nil
		default:
			return nil, os.ErrInvalid
		}
	}

	log.Info(logPrefix + " - request a source")
	start := time.Now()
	data, err := o.source.List(filter, paging)
	if err != nil {
		o.Store(key, err)

		log.Error(logPrefix+" - failed to request a source, ", err)
		return nil, err
	}
	log.Debug(logPrefix+" - request a source in ", time.Since(start))

	o.Store(key, data)

	return data, nil
}
