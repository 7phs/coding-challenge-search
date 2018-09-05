package model

import (
	log "github.com/sirupsen/logrus"
)

var (
	SearchModel *Search
)

type Dependencies struct {
	SearchDataSource SearchDataSource
	Lem              Lemmer
}

func Init(dep Dependencies) {
	log.Info("model: init")

	log.Info("model: create a model 'search'")
	SearchModel = NewSearch(dep.SearchDataSource)

	log.Info("model: make a factory keywords parser")
	NewSearchKeywords = newSearchKeywords(dep.Lem)
}
