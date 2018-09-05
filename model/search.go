package model

import (
	"bytes"
	"crypto/md5"
	"errors"
	"strings"
)

const (
	KeywordsModeDefault KeywordsMode = 0
	KeywordsModeAnd     KeywordsMode = 1
	KeywordsModeOr      KeywordsMode = 2
)

type KeywordsMode int

func (o KeywordsMode) String() string {
	switch o {
	case KeywordsModeAnd:
		return "AND"
	case KeywordsModeOr:
		return "OR"
	default:
		return "AND"
	}
}

type Lemmer interface {
	Parse(string) LemResult
}

type LemResult interface {
	Words() []string
	Lemmas() []string
}

type SearchFilter struct {
	Mode     KeywordsMode   `json:"searchTerms,omitempty"`
	Keywords *SearchKeyword `json:"searchTerms"`
	Location Location       `json:"location"`
}

func (o SearchFilter) Empty() bool {
	return o.Keywords.Empty() && o.Location.Empty()
}

func (o SearchFilter) String() string {
	return o.Mode.String() + "('" + o.Keywords.String() + "') + (" + o.Location.String() + ")"
}

func (o SearchFilter) Hash(b []byte) []byte {
	buf := bytes.NewBufferString(o.Mode.String())
	buf.WriteString(";")
	buf.WriteString(strings.Join(o.Keywords.Lemmas(), ">|<"))
	buf.WriteString(";")
	buf.WriteString(o.Location.String())

	hash := md5.New()

	if len(b) > 0 {
		hash.Write(b)
	}

	return hash.Sum(buf.Bytes())
}

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
