package memory

import (
	"os"
	"testing"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/db/memory/index"
	"github.com/7phs/coding-challenge-search/model"
	"github.com/7phs/coding-challenge-search/nlp"
	"github.com/stretchr/testify/assert"
)

const (
	modeMockIndexKeyword  = 1
	modeMockIndexLocation = 2
)

type mockIndex struct {
	index *index.ItemResult

	mode int
	rate float64
}

func newMockIndex(mode int, rate float64) *mockIndex {
	return &mockIndex{
		index: &index.ItemResult{},
		mode:  mode,
		rate:  rate,
	}
}

func (o *mockIndex) Add(record *model.Item) {
	o.index.Add(record, o.rate)

	o.rate -= 1e-4
}

func (o *mockIndex) Finish() {
	o.index.Sort()
}

func (o *mockIndex) Search(filter *model.SearchFilter) (index.Result, error) {
	switch o.mode {
	case modeMockIndexKeyword:
		if filter.Keywords.Empty() {
			return nil, os.ErrInvalid
		}

	case modeMockIndexLocation:
		if filter.Location.Empty() {
			return nil, os.ErrInvalid
		}
	}

	return o.index, nil
}

func TestItems_List(t *testing.T) {
	lemmer := nlp.NewMockLemmer()
	newSearchKeyword := model.FactorySearchKeywords(&config.Config{
		KeywordsLimit: 120,
	}, lemmer)

	data := model.NewMockSearchDataSource(100)

	items := NewItems(data,
		newMockIndex(modeMockIndexKeyword, 100000),
		newMockIndex(modeMockIndexLocation, 1000))

	err := items.Init()
	assert.Nil(t, err)

	count := 20
	exist, err := items.List(&model.SearchFilter{
		Keywords: newSearchKeyword("hello world"),
	}, &model.Paging{
		Limit: count,
	})
	assert.Nil(t, err)
	assert.Len(t, exist, count)

	exist, err = items.List(&model.SearchFilter{
		Location: model.Location{Lat: 10, Long: 10},
	}, &model.Paging{
		Limit: count,
	})
	assert.Nil(t, err)
	assert.Len(t, exist, count)

	exist, err = items.List(&model.SearchFilter{
		Keywords: newSearchKeyword("hello world"),
		Location: model.Location{Lat: 10, Long: 10},
	}, &model.Paging{
		Limit: count,
	})
	assert.Nil(t, err)
	assert.Len(t, exist, count)
}
