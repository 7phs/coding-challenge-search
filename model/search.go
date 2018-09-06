package model

import (
	"errors"
)

type SearchDataSource interface {
	List(*SearchFilter, *Paging) (ItemsList, error)
}

type Search struct {
	dataSource SearchDataSource
}

func NewSearch(dataSource SearchDataSource) *Search {
	return &Search{
		dataSource: dataSource,
	}
}

func (o *Search) List(filter *SearchFilter, paging *Paging) (ItemsList, error) {
	if filter.Empty() {
		return nil, errors.New("model/search: filter - empty")
	}

	result, err := o.dataSource.List(filter, paging)
	if err != nil {
		return nil, errors.New("model/search: failed to get a result for " + filter.String() + ", " + err.Error())
	}

	return result, nil
}
