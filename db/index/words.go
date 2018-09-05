package index

import (
	"os"

	"github.com/7phs/coding-challenge-search/model"
)

type LemmaStat struct {
	Stat map[string]int
}

func NewLemmaStat() *LemmaStat {
	return &LemmaStat{
		Stat: make(map[string]int),
	}
}

func (o *LemmaStat) Add(lemmas []string) *LemmaStat {
	for _, lemma := range lemmas {
		o.Stat[lemma]++
	}

	return o
}

type Words struct {
	lemmer model.Lemmer

	total   *LemmaStat
	mapping map[*model.Item]*LemmaStat
	index   map[string]*ItemResult
}

func NewIndexWords(lemmer model.Lemmer) *Words {
	return &Words{
		lemmer: lemmer,

		total:   NewLemmaStat(),
		mapping: make(map[*model.Item]*LemmaStat),
		index:   make(map[string]*ItemResult),
	}
}

func (o *Words) Add(record *model.Item) {
	lemmas := o.lemmer.Parse(record.Name).Lemmas()

	o.mapping[record] = NewLemmaStat().Add(lemmas)
	o.total.Add(lemmas)
}

func (o *Words) Finish() {
	for record, stat := range o.mapping {
		o.reindex(record, stat)
	}

	for _, res := range o.index {
		res.Sort()
	}
}

func (o *Words) reindex(record *model.Item, stat *LemmaStat) {
	var (
		result *ItemResult
		ok     bool
	)

	for lemma, stat := range stat.Stat {
		if result, ok = o.index[lemma]; !ok {
			result = &ItemResult{}

			o.index[lemma] = result
		}

		result.Add(record, o.calcRate(lemma, stat))
	}
}

func (o *Words) calcRate(lemma string, stat int) float64 {
	if v, ok := o.total.Stat[lemma]; ok && v > 0 {
		return float64(stat) / float64(v)
	}

	return .0
}

func (o *Words) Search(filter *model.SearchFilter) (Result, error) {
	if filter.Keywords.Empty() {
		return nil, os.ErrInvalid
	}

	var (
		total Result
		first = true
	)

	total = &ItemResult{}

	total.SetMode(filter.Mode)

	for _, lemma := range filter.Keywords.Lemmas() {
		if result, ok := o.index[lemma]; !ok {
			continue
		} else {
			if first {
				total.Append(result)
				first = !first
			} else {
				total = total.Reduce(result)
			}
		}
	}

	if total.Empty() {
		return nil, os.ErrNotExist
	}

	return total, nil
}
