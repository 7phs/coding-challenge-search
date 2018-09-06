package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	epsilon = 1e-5
)

func TestLocation_Validate(t *testing.T) {
	testSuites := []*struct {
		in      *Location
		empty   bool
		str     string
		isValid bool
	}{
		{
			in:      &Location{},
			empty:   true,
			str:     "lat: 0; long: 0",
			isValid: true,
		},
		{
			in:      &Location{Lat: 5, Long: -5},
			str:     "lat: 5; long: -5",
			isValid: true,
		},
		{
			in: &Location{Lat: 90.1, Long: -5},
		},
		{
			in: &Location{Lat: -90.1, Long: -5},
		},
		{
			in: &Location{Lat: 5, Long: -180.01},
		},
		{
			in: &Location{Lat: 5, Long: 180.01},
		},
		{
			in: &Location{Lat: 90.1, Long: 180.01},
		},
	}

	for i, test := range testSuites {
		assert.Equal(t, test.empty, test.in.Empty(), "%d: empty", i+1)
		if test.isValid {
			assert.Nil(t, test.in.Validate(), "%d: validation", i+1)
			assert.Equal(t, test.str, test.in.String(), "%d: string", i+1)
		} else {
			assert.NotNil(t, test.in.Validate(), "%d: validation", i+1)
		}
	}
}

func TestLocation_Distance(t *testing.T) {
	testSuites := []*struct {
		in1      Location
		in2      Location
		expected float64
	}{
		{
			in1:      Location{Lat: 55.751631, Long: 48.752316},
			in2:      Location{Lat: 55.820423, Long: 49.094051},
			expected: 22.694370,
		},
		{
			in1:      Location{Lat: 51.522634, Long: -0.084255},
			in2:      Location{Lat: 51.523029, Long: -0.084266},
			expected: 0.0439286,
		},
	}

	for i, test := range testSuites {
		test.in1.PreCalc()
		test.in2.PreCalc()

		exist := test.in1.Distance(test.in2)

		assert.InEpsilon(t, test.expected, exist, epsilon, "%d: distance", i+1)
	}
}

func TestLocationInt64_Compare(t *testing.T) {
	precision := 100
	testSuites := []*struct {
		in1      Location
		in2      LocationInt64
		expected int
	}{
		{
			in1:      Location{Lat: 55.751631, Long: 48.752316},
			in2:      LocationInt64{Lat: 5575, Long: 4875},
			expected: 0,
		},
		{
			in1:      Location{Lat: 55.751631, Long: 48.752316},
			in2:      LocationInt64{Lat: 5675, Long: 4875},
			expected: -1,
		},
		{
			in1:      Location{Lat: 55.751631, Long: 48.752316},
			in2:      LocationInt64{Lat: 5575, Long: 4975},
			expected: -1,
		},
		{
			in1:      Location{Lat: 55.751631, Long: 48.752316},
			in2:      LocationInt64{Lat: 5475, Long: 4975},
			expected: 1,
		},
		{
			in1:      Location{Lat: 55.751631, Long: 48.752316},
			in2:      LocationInt64{Lat: 5575, Long: 4775},
			expected: 1,
		},
	}

	for i, test := range testSuites {
		exist := NewLocationInt64(test.in1, precision)
		assert.Equal(t, test.expected, exist.Compare(&test.in2), "%d: compare", i+1)
	}
}

func TestLocationInt64_Distance(t *testing.T) {
	testSuites := []*struct {
		in1       Location
		precision int
		expected  LocationInt64
		in2       Location
		expDist   float64
	}{
		{
			in1:       Location{Lat: 55.751631, Long: 48.752316},
			precision: 100,
			expected:  LocationInt64{Lat: 5575, Long: 4875, loc: Location{Lat: 55.75, Long: 48.75}},

			in2:     Location{Lat: 55.820423, Long: 49.094051},
			expDist: 22.892577,
		},
		{
			in1:       Location{Lat: 55.751631, Long: 48.752316},
			precision: 1,
			expected:  LocationInt64{Lat: 55, Long: 48, loc: Location{Lat: 55., Long: 48.}},
			in2:       Location{Lat: 55.820423, Long: 49.094051},
			expDist:   114.416979,
		},
	}

	for i, test := range testSuites {
		exist := NewLocationInt64(test.in1, test.precision)
		if assert.Equal(t, &test.expected, exist) {
			exist.PreCalc()
			test.in2.PreCalc()

			existDist := exist.Distance(test.in2)

			assert.InEpsilon(t, test.expDist, existDist, epsilon, "%d: distance", i+1)
		}

	}

}
