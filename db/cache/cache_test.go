package cache

import (
	"testing"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/model"
	"github.com/7phs/coding-challenge-search/nlp"
	"github.com/stretchr/testify/assert"
)

func TestItems_List(t *testing.T) {
	var (
		mockDataSource    = model.NewMockSearchDataSource(100)
		newSearchKeywords = model.FactorySearchKeywords(&config.Config{
			KeywordsLimit: 200,
		}, nlp.NewMockLemmer())
		itemsCache  = NewItems(mockDataSource)
		repeatCount = 4
	)

	filterPaging := []*struct {
		filter *model.SearchFilter
		paging *model.Paging
		err    bool
	}{
		{
			paging: &model.Paging{Start: 0, Limit: 10},
		},
		{
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeOr,
				Keywords: newSearchKeywords("hello worlds"),
				Location: model.Location{Lat: 10., Long: 15.},
			},
			paging: &model.Paging{Start: 10, Limit: 20},
		},
		{ // the same as a previous, but different points of filter and paging
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeOr,
				Keywords: newSearchKeywords("hello worlds"),
				Location: model.Location{Lat: 10., Long: 15.},
			},
			paging: &model.Paging{Start: 10, Limit: 20},
		},
		{
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeAnd,
				Keywords: newSearchKeywords("hello worlds"),
				Location: model.Location{Lat: 10., Long: 15.},
			},
			paging: &model.Paging{Start: 20, Limit: 30},
		},
		{
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeAnd,
				Keywords: newSearchKeywords("hello worlds"),
				Location: model.Location{Lat: 10., Long: 15.},
			},
			paging: &model.Paging{Start: len(mockDataSource.Items) + 1, Limit: 30},
			err:    true,
		},
	}

	for i, suite := range filterPaging {
		for j := 0; j < repeatCount; j++ {
			list, err := itemsCache.List(suite.filter, suite.paging)
			if suite.err {
				assert.NotNil(t, err, "%d: request", i+1)
			} else if assert.Nil(t, err, "%d: request", i+1) {

				start, limit, err := suite.paging.StartLimit(len(mockDataSource.Items))
				assert.Nil(t, err, "%d: start/limit ", i+1)

				assert.Equal(t, mockDataSource.Items[start:limit], list, "%d: list", i+1)
			}
		}
	}

	assert.Equal(t, 4, mockDataSource.CallCount, "request for data source")
}
