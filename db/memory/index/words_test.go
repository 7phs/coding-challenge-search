package index

import (
	"strings"
	"testing"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/model"
	"github.com/7phs/coding-challenge-search/nlp"
	"github.com/stretchr/testify/assert"
)

func TestLemmaStat_Add(t *testing.T) {
	exist := NewLemmaStat()

	expected := map[string]int{
		"hello":  15,
		"world":  10,
		"привет": 3,
	}
	for word, count := range expected {
		lemmas := make([]string, 0, count)
		for i := 0; i < count; i++ {
			lemmas = append(lemmas, word)
		}

		exist.Add(lemmas)
	}

	assert.Equal(t, expected, exist.Stat)
}

func TestWords_Search(t *testing.T) {
	lemmer := nlp.NewMockLemmer()
	idx := NewIndexWords(lemmer)
	newSearchKeyword := model.FactorySearchKeywords(&config.Config{
		KeywordsLimit: 120,
	}, lemmer)

	data := model.ItemsList{
		// 0
		&model.Item{
			Id:   1,
			Name: strings.Repeat("word1 ", 5) + strings.Repeat("word2 ", 3),
		},
		// 1
		&model.Item{
			Id:   2,
			Name: strings.Repeat("word2 ", 5) + strings.Repeat("word3 ", 3),
		},
		// 2
		&model.Item{
			Id:   3,
			Name: strings.Repeat("word3 ", 5) + strings.Repeat("word4 ", 3),
		},
		// 3
		&model.Item{
			Id: 4,
			Name: strings.Repeat("word1 ", 4) +
				strings.Repeat("word2 ", 4) +
				strings.Repeat("word3 ", 4) +
				strings.Repeat("word4 ", 4),
		},
	}

	for _, rec := range data {
		idx.Add(rec)
	}

	idx.Finish()

	paging := &model.Paging{Limit: 20}

	testSuites := []*struct {
		filter   *model.SearchFilter
		err      bool
		expected model.ItemsList
	}{
		{
			filter: &model.SearchFilter{},
			err:    true,
		},
		{
			filter: &model.SearchFilter{
				Keywords: newSearchKeyword("word1"),
			},
			expected: model.ItemsList{data[0], data[3]},
		},
		{
			filter: &model.SearchFilter{
				Keywords: newSearchKeyword("word2"),
			},
			expected: model.ItemsList{data[1], data[3], data[0]},
		},
		{
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeAnd,
				Keywords: newSearchKeyword("word1 word3"),
			},
			expected: model.ItemsList{data[3]},
		},
		{
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeOr,
				Keywords: newSearchKeyword("word1 word3"),
			},
			expected: model.ItemsList{data[3], data[0], data[2], data[1]},
		},
		{
			filter: &model.SearchFilter{
				Mode:     model.KeywordsModeOr,
				Keywords: newSearchKeyword("unknown"),
			},
			err: true,
		},
	}

	for _, test := range testSuites {
		exist, err := idx.Search(test.filter)
		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)

			assert.Equal(t, test.expected, exist.Items(paging))
		}
	}
}
