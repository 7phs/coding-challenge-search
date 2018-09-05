package model

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/errCode"
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

func keywordsFilter(keywords string) (string, error) {
	var err error

	keywords = strings.TrimSpace(strings.ToLower(keywords))
	if len(keywords) > int(config.Conf.KeywordsLimit) {
		err = errors.New("k: length is greater than allowed " + config.Conf.KeywordsLimit.String())
	} else {
		keywords = filterAlphaNum.ReplaceAllString(keywords, " ")
	}

	return keywords, err
}

func newSearchKeywords(lemmer Lemmer) SearchKeywordFactory {
	return func(keywords string) *SearchKeyword {
		o := &SearchKeyword{}

		keywords, o.err = keywordsFilter(keywords)
		if o.err == nil {
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
	return o.words
}

func (o *SearchKeyword) Lemmas() []string {
	return o.lemmas
}

func (o *SearchKeyword) String() string {
	return strings.Join(o.words, ", ")
}

func (o *SearchKeyword) Validate() (errList errCode.ErrList) {
	if o.err != nil {
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
