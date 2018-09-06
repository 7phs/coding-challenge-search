package model

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/errCode"
	"github.com/7phs/coding-challenge-search/nlp"
	"github.com/c2h5oh/datasize"
)

const (
	limitWordsCount = 32
)

var (
	filterAlphaNum    = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9\-]+`)
	NewSearchKeywords SearchKeywordFactory
)

type SearchKeywordFactory func(string) *SearchKeyword

type SearchKeyword struct {
	words  []string
	lemmas []string
	err    error
}

func keywordsFilter(keywords string, limit int) (string, error) {
	var err error

	keywords = strings.TrimSpace(strings.ToLower(keywords))
	if len(keywords) > limit {
		err = errors.New("k: length is greater than allowed " + strconv.Itoa(limit))
	} else {
		keywords = filterAlphaNum.ReplaceAllString(keywords, " ")
	}

	return keywords, err
}

func FactorySearchKeywords(conf *config.Config, lemmer nlp.Lemmer) SearchKeywordFactory {
	return func(keywords string) *SearchKeyword {
		o := &SearchKeyword{}

		keywords, o.err = keywordsFilter(keywords, int(conf.KeywordsLimit))
		if o.err == nil && len(keywords) > 0 {
			res := lemmer.Parse(keywords)

			o.words = res.Words()
			o.lemmas = res.Lemmas()
		}

		return o
	}
}

func (o *SearchKeyword) Empty() bool {
	return o == nil || len(o.words) == 0
}

func (o *SearchKeyword) Words() []string {
	if o == nil {
		return nil
	}

	return o.words
}

func (o *SearchKeyword) Lemmas() []string {
	if o == nil {
		return nil
	}

	return o.lemmas
}

func (o *SearchKeyword) String() string {
	if o == nil {
		return "nil"
	}

	return strings.Join(o.words, ", ")
}

func (o *SearchKeyword) Validate() (errList errCode.ErrList) {
	if o == nil {
		errList.Add(errors.New("empty"))
	} else if o.err != nil {
		errList.Add(o.err)
	} else {
		if len(o.words) == 0 {
			errList.Add(errors.New("k: empty"))
		} else if len(o.words) > limitWordsCount {
			errList.Add(fmt.Errorf("k: words count is greater than allowed %d", limitWordsCount))
		}
	}

	return
}

func MockNewSearchKeywords(limit int) func() {
	preNewSearchKeywords := NewSearchKeywords

	NewSearchKeywords = FactorySearchKeywords(&config.Config{
		KeywordsLimit: datasize.ByteSize(limit),
	}, nlp.NewMockLemmer())

	return func() {
		NewSearchKeywords = preNewSearchKeywords
	}
}
