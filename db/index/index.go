package index

import (
	"time"

	"github.com/7phs/coding-challenge-search/model"
	log "github.com/sirupsen/logrus"
)

type ItemIndex interface {
	Add(record *model.Item)
	Finish()

	Search(filter *model.SearchFilter) (Result, error)
}

type Result interface {
	SetMode(model.KeywordsMode)
	Append(Result)
	Reduce(Result) Result
	Empty() bool

	Items(int, int) model.ItemsList
	ItemsWithRateById() model.ItemWithRateList
	ItemsWithRateByRate() model.ItemWithRateList
}

type ItemResult struct {
	mode          model.KeywordsMode
	recordsById   model.ItemWithRateList
	recordsByRate model.ItemWithRateList
}

func (o *ItemResult) Add(record *model.Item, rate float64) {
	ratedRecord := model.NewItemWithRate(record, rate)

	o.recordsById.Add(ratedRecord)
	o.recordsByRate.Add(ratedRecord)
}

func (o *ItemResult) Normalize(maxDistance float64) {
	for _, item := range o.recordsById {
		item.Rate = 1 - item.Rate/maxDistance
	}
}

func (o *ItemResult) Sort() {
	o.recordsById.SortById()
	o.recordsByRate.Sort()
}

func (o *ItemResult) SetMode(mode model.KeywordsMode) {
	o.mode = mode
}

func (o *ItemResult) Append(result Result) {
	if items := result.ItemsWithRateByRate(); items != nil {
		o.recordsByRate = append(o.recordsByRate, items...)
	}

	o.recordsById = append(o.recordsById, result.ItemsWithRateById()...)
}

func (o *ItemResult) Reduce(result Result) Result {
	var records model.ItemWithRateList

	switch o.mode {
	case model.KeywordsModeAnd:
		records = o.recordsById.Intersect(result.ItemsWithRateById())
	case model.KeywordsModeOr:
		records = o.recordsById.Join(result.ItemsWithRateById())
	default:
		records = o.recordsById.Intersect(result.ItemsWithRateById())
	}

	return &ItemResult{
		mode:        o.mode,
		recordsById: records,
	}
}

func (o *ItemResult) Empty() bool {
	return o == nil || len(o.recordsById) == 0
}

func (o *ItemResult) Items(start, limit int) model.ItemsList {
	// specific case - no one joining or intersect
	// just using already sorted list
	if o.recordsByRate != nil {
		if start > len(o.recordsByRate) {
			return nil
		}

		if start+limit > len(o.recordsByRate) {
			limit = len(o.recordsByRate) - start
		}

		return o.recordsByRate[start:limit].ItemsList()
	}

	s := time.Now()
	copied := o.recordsById.Copy()
	log.Debug("copied ", time.Since(s))

	s = time.Now()
	copied.Sort()
	log.Debug("sorted ", time.Since(s))

	if start > len(copied) {
		return nil
	}

	if start+limit > len(copied) {
		limit = len(copied) - start
	}

	copied = copied[start:limit]

	return copied.ItemsList()
}

func (o *ItemResult) ItemsWithRateById() model.ItemWithRateList {
	return o.recordsById
}

func (o *ItemResult) ItemsWithRateByRate() model.ItemWithRateList {
	return o.recordsByRate
}
