package model

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaging_Hash(t *testing.T) {
	testSuites := []*struct {
		in       *Paging
		ext      []byte
		expected string
	}{
		{
			in:       &Paging{Start: 3, Limit: 56},
			expected: "333b3536d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			in:       &Paging{Start: -3, Limit: -56},
			ext:      []byte{34, 56, 170},
			expected: "2d333b2d353602bc4739fb7a1f01cecadb1438642dc0",
		},
	}

	for _, test := range testSuites {
		exist := test.in.Hash(test.ext)
		expected, _ := hex.DecodeString(test.expected)

		assert.Equal(t, []byte(expected), exist)
	}
}

func TestPaging_StartLimit(t *testing.T) {
	testSuites := []*struct {
		in       *Paging
		ln       int
		expStart int
		expLimit int
		expErr   bool
	}{
		{
			in:       &Paging{Start: 3, Limit: 56},
			ln:       127,
			expStart: 3,
			expLimit: 59,
		},
		{
			in:     &Paging{Start: 50, Limit: 6},
			ln:     40,
			expErr: true,
		},
		{
			in:       &Paging{Start: 30, Limit: 20},
			ln:       40,
			expStart: 30,
			expLimit: 40,
		},
	}

	for i, test := range testSuites {
		existStart, existLimit, err := test.in.StartLimit(test.ln)

		if test.expErr {
			assert.NotNil(t, err, "%d: checking error", i+1)
		} else if assert.Nil(t, err) {
			assert.Equal(t, test.expStart, existStart, "%d: start", i+1)
			assert.Equal(t, test.expLimit, existLimit, "%d: limit", i+1)
		}
	}
}
