package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullString(t *testing.T) {
	err := (&NullString{}).Scan("hello")
	assert.Nil(t, err)

	testSuites := []*struct {
		exist    string
		existNil bool
		value    interface{}
		expected string
	}{
		{
			exist:    "not-set",
			value:    nil,
			expected: "",
		},
		{
			exist:    "",
			existNil: true,
			value:    "1456",
			expected: "",
		},
		{
			exist:    "",
			value:    "1456",
			expected: "1456",
		},
	}

	for i, test := range testSuites {
		nullV := &NullString{}
		if !test.existNil {
			nullV.V = &test.exist
		}

		err = nullV.Scan(test.value)
		assert.Nil(t, err)
		assert.Equal(t, test.expected, test.exist, "%d: value", i+1)
	}

	v, err := NullString{}.Value()
	assert.Nil(t, v, "value")
	assert.Nil(t, err, "err")

	exist := "1234"
	v, err = NullString{V: &exist}.Value()
	assert.Equal(t, exist, v, "value")
	assert.Nil(t, err, "err")
}

func TestNullFloat64(t *testing.T) {
	err := (&NullFloat64{}).Scan("hello")
	assert.Nil(t, err)

	err = (&NullFloat64{}).Scan("123")
	assert.Nil(t, err)

	testSuites := []*struct {
		exist    float64
		existNil bool
		value    interface{}
		expected float64
		err      bool
	}{
		{
			exist:    12345.12,
			value:    nil,
			expected: 0,
		},
		{
			exist:    0,
			existNil: true,
			value:    1456.12,
			expected: 0,
		},
		{
			exist:    0,
			value:    "1456.732",
			expected: 1456.732,
		},
		{
			exist:    9876.43,
			value:    "a1456",
			expected: 9876.43,
			err:      true,
		},
	}

	for i, test := range testSuites {
		nullV := &NullFloat64{}
		if !test.existNil {
			nullV.V = &test.exist
		}

		err = nullV.Scan(test.value)
		if test.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, test.expected, test.exist, "%d: value", i+1)
	}

	v, err := NullFloat64{}.Value()
	assert.Nil(t, v, "value")
	assert.Nil(t, err, "err")

	exist := float64(1234.954)
	v, err = NullFloat64{V: &exist}.Value()
	assert.Equal(t, exist, v, "value")
	assert.Nil(t, err, "err")
}
