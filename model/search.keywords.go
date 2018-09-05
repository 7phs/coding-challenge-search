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
	filterAlphaNum = regexp.MustCompile(`[^a-zA-Zа-яА-Я0-9\-]+`)
)

type SearchKeyword struct {
	words []string
	err   error
}

func NewSearchKeywords(keywords string) *SearchKeyword {
	var (
		words = make([]string, 0)
		err   error
	)

	keywords = strings.TrimSpace(strings.ToLower(keywords))
	if len(keywords) > int(config.Conf.KeywordsLimit) {
		err = errors.New("k: length is greater than allowed " + config.Conf.KeywordsLimit.String())
	} else {
		keywords = filterAlphaNum.ReplaceAllString(keywords, " ")

		for _, word := range strings.Split(keywords, " ") {
			word := strings.TrimSpace(strings.ToLower(word))
			if len(word) == 0 {
				continue
			}

			words = append(words, word)
		}
	}

	return &SearchKeyword{
		words: words,
		err:   err,
	}
}

func (o *SearchKeyword) Empty() bool {
	return o == nil || len(o.words) == 0
}

func (o *SearchKeyword) Words() []string {
	return o.words
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
