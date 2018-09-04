package model

import (
	"errors"
	"fmt"

	"github.com/7phs/coding-challenge-search/errCode"
)

type Location struct {
	Lat  float64 `json:"lat" form:"lat"`
	Long float64 `json:"long" form:"long"`
}

func (o Location) String() string {
	return fmt.Sprint("lat: ", o.Lat, "; long: ", o.Long)
}

func (o Location) Validate() error {
	var errList errCode.ErrList

	if o.Lat < -90. || o.Lat > 90. {
		errList.Add(errors.New("lat: out of bounds"))
	}

	if o.Long < -180. || o.Long > 180. {
		errList.Add(errors.New("long: out of bounds"))
	}

	return errList.Result()
}

func (o Location) Empty() bool {
	return o.Lat == 0 && o.Long == 0
}

type SearchRecord struct {
	Title    string   `json:"title"`
	Location Location `json:"loc"`
	Url      struct {
		Item string   `json:"item"`
		Imgs []string `json:"imgs"`
	} `json:"url"`
}

type SearchResult []*SearchRecord
