package model

import (
	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/nlp"
	log "github.com/sirupsen/logrus"
)

var (
	SearchModel *Search
)

type Dependencies struct {
	SearchDataSource SearchDataSource
	Lem              nlp.Lemmer
}

func Init(dep Dependencies) {
	log.Info("model: init")

	log.Info("model: create a model 'search'")
	SearchModel = NewSearch(dep.SearchDataSource)

	log.Info("model: make a factory keywords parser")
	NewSearchKeywords = FactorySearchKeywords(config.Conf, dep.Lem)
}
