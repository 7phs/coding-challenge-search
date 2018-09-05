package model

import (
	log "github.com/sirupsen/logrus"
)

var (
	SearchModel *Search
)

type Dependencies struct {
	SearchDataSource SearchDataSource
}

func Init(dep Dependencies) {
	log.Info("model: init")

	log.Info("model: create a model 'search'")
	SearchModel = NewSearch(dep.SearchDataSource)
}
