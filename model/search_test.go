package model

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchFilter_String(t *testing.T) {
	defer MockNewSearchKeywords(200)()

	testSuites := []*struct {
		in      *SearchFilter
		empty   bool
		str     string
		extHash []byte
		hash    string
	}{
		{
			empty: true,
			str:   "nil",
		},
		{
			in: &SearchFilter{
				Mode:     KeywordsModeDefault,
				Keywords: NewSearchKeywords("hello world"),
				Location: Location{Lat: 10.5, Long: -13.6},
			},
			str:  "AND('hello, world') + (lat: 10.5; long: -13.6)",
			hash: "414e443b68656c6c6f3e7c3c776f726c643b6c61743a2031302e353b206c6f6e673a202d31332e36d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			in: &SearchFilter{
				Mode:     KeywordsModeAnd,
				Location: Location{Lat: 10.5, Long: -13.6},
			},
			str:  "AND('nil') + (lat: 10.5; long: -13.6)",
			hash: "414e443b3b6c61743a2031302e353b206c6f6e673a202d31332e36d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			in: &SearchFilter{
				Mode:     KeywordsModeOr,
				Keywords: NewSearchKeywords("привет"),
			},
			str:     "OR('привет') + (lat: 0; long: 0)",
			extHash: []byte{45, 129, 92},
			hash:    "4f523bd0bfd180d0b8d0b2d0b5d1823b6c61743a20303b206c6f6e673a2030851c8379c44982c088b1c449a75fff46",
		},
	}

	for i, test := range testSuites {
		var hash []byte
		if len(test.hash) == 0 {
			hash = nil
		} else {
			hash, _ = hex.DecodeString(test.hash)
		}

		assert.Equal(t, test.empty, test.in.Empty(), "%d: empty", i+1)
		assert.Equal(t, test.str, test.in.String(), "%d: string", i+1)
		assert.Equal(t, hash, test.in.Hash(test.extHash), "%d: hash", i+1)
	}
}

func TestSearchKeyword_Validate(t *testing.T) {
	defer MockNewSearchKeywords(70)()

	testSuites := []*struct {
		in     *SearchKeyword
		empty  bool
		str    string
		words  []string
		lemmas []string
		err    bool
	}{
		{
			empty: true,
			str:   "nil",
			err:   true,
		},
		{
			in:     NewSearchKeywords("hello world"),
			str:    "hello, world",
			words:  []string{"hello", "world"},
			lemmas: []string{"hello", "world"},
		},
		{
			in:    NewSearchKeywords(`hello world as the longest line! hello world as the longest line! hello`),
			empty: true,
			err:   true,
		},
		{
			in:     NewSearchKeywords(`h e l l o w o r l d a s t h e l o n g e s t h e l l o w o r l d a s`),
			str:    "h, e, l, l, o, w, o, r, l, d, a, s, t, h, e, l, o, n, g, e, s, t, h, e, l, l, o, w, o, r, l, d, a, s",
			words:  []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d", "a", "s", "t", "h", "e", "l", "o", "n", "g", "e", "s", "t", "h", "e", "l", "l", "o", "w", "o", "r", "l", "d", "a", "s"},
			lemmas: []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d", "a", "s", "t", "h", "e", "l", "o", "n", "g", "e", "s", "t", "h", "e", "l", "l", "o", "w", "o", "r", "l", "d", "a", "s"},
			err:    true,
		},
		{
			in:    NewSearchKeywords(""),
			empty: true,
			err:   true,
		},
	}

	for i, test := range testSuites {
		if test.err {
			assert.NotNil(t, test.in.Validate(), "%d: validate", i+1)
		} else {
			assert.Nil(t, test.in.Validate(), "%d: validate", i+1)
		}

		assert.Equal(t, test.empty, test.in.Empty(), "%d: empty", i+1)
		assert.Equal(t, test.str, test.in.String(), "%d: string", i+1)
		assert.Equal(t, test.words, test.in.Words(), "%d: words", i+1)
		assert.Equal(t, test.lemmas, test.in.Lemmas(), "%d: lemmas", i+1)
	}
}

func TestSearch_List(t *testing.T) {
	defer MockNewSearchKeywords(200)()

	var (
		count          = 100
		mockDataSource = NewMockSearchDataSource(count)
		searchModel    = NewSearch(mockDataSource)
		paging         = &Paging{
			Start: 5,
			Limit: 10,
		}
	)
	// EMPTY FILTER
	_, err := searchModel.List(nil, paging)
	assert.NotNil(t, err, "empty filter")
	// A GENERAL
	exist, err := searchModel.List(&SearchFilter{
		Keywords: NewSearchKeywords("hello world"),
	}, paging)
	assert.Nil(t, err)
	assert.Equal(t, mockDataSource.Items[5:15], exist)
	// A DATA SOURCE ERROR
	_, err = searchModel.List(&SearchFilter{
		Keywords: NewSearchKeywords("hello world"),
	}, &Paging{
		Start: count + 5,
		Limit: 10,
	})
	assert.NotNil(t, err)
}
