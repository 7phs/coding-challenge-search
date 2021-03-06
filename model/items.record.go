package model

import (
	"encoding/json"
	"sort"
)

type Item struct {
	Id       int64    `json:"Id"`
	Name     string   `json:"item_name"`
	Location Location `json:"loc"`
	Url      string   `json:"item_url"`
	Imgs     []string `json:"img_urls"`
}

func (o *Item) String() string {
	d, _ := json.Marshal(o)

	return string(d)
}

type ItemsList []*Item

func (o *ItemsList) Add(item *Item) {
	*o = append(*o, item)
}

func (o *ItemsList) Sort() {
	sort.Slice(*o, func(i, j int) bool {
		return (*o)[i].Id < (*o)[j].Id
	})
}

func (o ItemsList) Has(rec *Item) bool {
	index := sort.Search(len(o), func(i int) bool {
		return o[i].Id >= rec.Id
	})

	return index < len(o) && o[index].Id == rec.Id
}

type ItemWithRate struct {
	*Item

	Rate float64
}

func (o *ItemWithRate) CompareId(r *ItemWithRate) int {
	if o.Id < r.Id {
		return -1
	} else if o.Id > r.Id {
		return 1
	}

	return 0
}

func NewItemWithRate(record *Item, rate float64) *ItemWithRate {
	return &ItemWithRate{
		Item: record,
		Rate: rate,
	}
}

type ItemWithRateList []*ItemWithRate

func (o ItemWithRateList) ItemsList() ItemsList {
	if len(o) == 0 {
		return nil
	}

	result := make(ItemsList, 0, len(o))

	for _, rec := range o {
		result.Add(rec.Item)
	}

	return result
}

func (o *ItemWithRateList) Add(record *ItemWithRate) {
	*o = append(*o, record)
}

func (o *ItemWithRateList) SortById() ItemWithRateList {
	sort.Slice(*o, func(i, j int) bool {
		return (*o)[i].CompareId((*o)[j]) < 0
	})

	return *o
}

func (o *ItemWithRateList) Sort() ItemWithRateList {
	sort.Slice(*o, func(i, j int) bool {
		return (*o)[i].Rate > (*o)[j].Rate
	})

	return *o
}

// Both lists should be sorted
func (o ItemWithRateList) Intersect(r ItemWithRateList) ItemWithRateList {
	if len(o) == 0 || len(r) == 0 {
		return nil
	}

	result := make(ItemWithRateList, 0, (len(o)+len(r))/2)

	for i, j := 0, 0; i < len(o) && j < len(r); {
		switch o[i].CompareId(r[j]) {
		case 0:
			o[i].Rate += r[j].Rate

			result.Add(o[i])
			i++
			j++

		case -1:
			i++

		case 1:
			j++
		}
	}

	return result
}

func (o ItemWithRateList) Join(r ItemWithRateList) ItemWithRateList {
	result := make(ItemWithRateList, 0, (len(o)+len(r))/2)
	i, j := 0, 0

	for i < len(o) && j < len(r) {
		switch o[i].CompareId(r[j]) {
		case 0:
			o[i].Rate += r[j].Rate

			result.Add(o[i])

			i++
			j++

		case -1:
			result.Add(o[i])
			i++

		case 1:
			result.Add(r[j])
			j++
		}
	}
	// DON'T FORGOT REST
	for ; i < len(o); i++ {
		result.Add(o[i])
	}

	for ; j < len(r); j++ {
		result.Add(r[j])
	}

	return result
}

func (o ItemWithRateList) Copy() ItemWithRateList {
	copied := make(ItemWithRateList, len(o))

	copy(copied, o)

	return copied
}
