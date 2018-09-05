package handler

import "github.com/7phs/coding-challenge-search/model"

type MetaResponse struct {
	Search *MetaSearchResponse `json:"search,omitempty"`
}

type MetaSearchResponse struct {
	Keywords string         `json:"keywords"`
	Location model.Location `json:"location"`
}
