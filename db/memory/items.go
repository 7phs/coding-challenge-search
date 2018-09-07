package memory

import (
	"errors"
	"os"
	"sync"
	"time"

	"github.com/7phs/coding-challenge-search/db/memory/index"
	"github.com/7phs/coding-challenge-search/model"
	log "github.com/sirupsen/logrus"
)

type ItemsSource interface {
	Load() (model.ItemsList, error)
}

type itemsListResult struct {
	result index.Result
	err    error
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
	log.Info(logPrefix + " reindexing records in " + time.Since(start).String())

	return nil
}

func (o *Items) load() (err error) {
	o.list, err = o.source.Load()

	return
}

func (o *Items) reindex() error {
	var wait sync.WaitGroup

	for _, idx := range o.indexes {
		wait.Add(1)
		go func(idx index.ItemIndex) {
			defer wait.Done()

			for _, record := range o.list {
				idx.Add(record)
			}

			idx.Finish()
		}(idx)
	}

	wait.Wait()

	return nil
}

func (o *Items) List(filter *model.SearchFilter, paging *model.Paging) (model.ItemsList, error) {
	var (
		total      index.Result
		wait       sync.WaitGroup
		waitResult sync.WaitGroup
		ch         = make(chan *itemsListResult)
	)

	for _, idx := range o.indexes {
		wait.Add(1)
		go func(idx index.ItemIndex) {
			defer wait.Done()

			result, err := idx.Search(filter)
			if err != nil {
				if err != os.ErrInvalid && err != os.ErrNotExist {
					log.Error("memory/items: failed to request a search result, ", err)
				}
			}

			ch <- &itemsListResult{
				result: result,
				err:    err,
			}
		}(idx)
	}

	waitResult.Add(1)
	go func() {
		defer waitResult.Done()

		for res := range ch {
			if res.err != nil || res.result == nil {
				continue
			}

			if total == nil {
				total = res.result
				total.SetMode(model.KeywordsModeAnd)
			} else {
				total = total.Reduce(res.result)
			}
		}
	}()

	wait.Wait()

	close(ch)
	waitResult.Wait()

	if total == nil {
		return nil, errors.New("memory/items: not found")
	}

	result := total.Items(paging)

	if result == nil {
		return nil, errors.New("memory/items: start is out of bound")
	}

	return result, nil
}
