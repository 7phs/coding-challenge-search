package handler

import "github.com/7phs/coding-challenge-search/model"

type MetaResponse struct {
	Search *MetaSearchResponse `json:"search,omitempty"`
}

type MetaSearchResponse struct {
	Filter *model.SearchFilter `json:"filter"`
	Paging *model.Paging       `json:"paging"`
}
