package common

import (
	"testing"
	"time"

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

func TestNullInt(t *testing.T) {
	err := (&NullInt{}).Scan("hello")
	assert.Nil(t, err)

	err = (&NullInt{}).Scan("123")
	assert.Nil(t, err)

	testSuites := []*struct {
		exist    int
		existNil bool
		value    interface{}
		expected int
		err      bool
	}{
		{
			exist:    12345,
			value:    nil,
			expected: 0,
		},
		{
			exist:    0,
			existNil: true,
			value:    1456,
			expected: 0,
		},
		{
			exist:    0,
			value:    "1456",
			expected: 1456,
		},
		{
			exist:    9876,
			value:    "a1456",
			expected: 9876,
			err:      true,
		},
	}

	for i, test := range testSuites {
		nullV := &NullInt{}
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

	v, err := NullInt{}.Value()
	assert.Nil(t, v, "value")
	assert.Nil(t, err, "err")

	exist := 1234
	v, err = NullInt{V: &exist}.Value()
	assert.Equal(t, exist, v, "value")
	assert.Nil(t, err, "err")
}

func TestNullInt64(t *testing.T) {
	err := (&NullInt64{}).Scan("hello")
	assert.Nil(t, err)

	err = (&NullInt64{}).Scan("123")
	assert.Nil(t, err)

	testSuites := []*struct {
		exist    int64
		existNil bool
		value    interface{}
		expected int64
		err      bool
	}{
		{
			exist:    12345,
			value:    nil,
			expected: 0,
		},
		{
			exist:    0,
			existNil: true,
			value:    1456,
			expected: 0,
		},
		{
			exist:    0,
			value:    "1456",
			expected: 1456,
		},
		{
			exist:    9876,
			value:    "a1456",
			expected: 9876,
			err:      true,
		},
	}

	for i, test := range testSuites {
		nullV := &NullInt64{}
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

	v, err := NullInt64{}.Value()
	assert.Nil(t, v, "value")
	assert.Nil(t, err, "err")

	exist := int64(1234)
	v, err = NullInt64{V: &exist}.Value()
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

func TestNullTime(t *testing.T) {
	err := (&NullTime{}).Scan("hello")
	assert.Nil(t, err)

	err = (&NullTime{}).Scan("123")
	assert.Nil(t, err)

	testSuites := []*struct {
		exist    time.Time
		existNil bool
		value    interface{}
		expected time.Time
		err      bool
	}{
		{
			exist:    time.Date(2018, 9, 8, 13, 56, 0, 0, time.UTC),
			value:    nil,
			expected: time.Time{},
		},
		{
			exist:    time.Time{},
			existNil: true,
			value:    time.Now(),
			expected: time.Time{},
		},
		{
			exist:    time.Time{},
			value:    "2018-09-08T13:56:00Z",
			expected: time.Date(2018, 9, 8, 13, 56, 0, 0, time.UTC),
		},
		{
			exist:    time.Time{},
			value:    time.Date(2018, 9, 8, 13, 56, 0, 0, time.UTC),
			expected: time.Date(2018, 9, 8, 13, 56, 0, 0, time.UTC),
		},
		{
			exist:    time.Date(2018, 9, 8, 13, 56, 0, 0, time.UTC),
			value:    "a1456",
			expected: time.Time{},
			err:      true,
		},
	}

	for i, test := range testSuites {
		nullV := &NullTime{}
		if !test.existNil {
			nullV.V = &test.exist
		}

		err = nullV.Scan(test.value)
		if test.err {
			assert.NotNil(t, err)
		} else if !assert.Nil(t, err, "%d: err", i+1) {
			t.Error(err.Error())
		}
		assert.Equal(t, test.expected, test.exist, "%d: value", i+1)
	}

	v, err := NullTime{}.Value()
	assert.Nil(t, v, "value")
	assert.Nil(t, err, "err")

	exist := time.Date(2018, 9, 8, 13, 56, 0, 0, time.UTC)
	v, err = NullTime{V: &exist}.Value()
	assert.Equal(t, exist, v, "value")
	assert.Nil(t, err, "err")
}
