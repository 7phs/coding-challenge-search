package model

import (
	"errors"
	"fmt"
	"math"

	"github.com/7phs/coding-challenge-search/errCode"
)

const (
	EarthMidRadius    = 6371.0088
	RadialCoefficient = math.Pi / 180.0
)

type LocationInt64 struct {
	loc Location

	Lat  int
	Long int
}

func NewLocationInt64(loc Location, precision int) *LocationInt64 {
	lat := int(loc.Lat * float64(precision))
	long := int(loc.Long * float64(precision))

	return &LocationInt64{
		loc: Location{
			Lat:  float64(lat / precision),
			Long: float64(long / precision),
		},

		Lat:  lat,
		Long: long,
	}
}

func (o *LocationInt64) PreCalc() *LocationInt64 {
	o.loc.PreCalc()

	return o
}

func (o *LocationInt64) Distance(r Location) float64 {
	return o.loc.Distance(r)
}

func (o *LocationInt64) Compare(r *LocationInt64) int {
	if o.Lat < r.Lat {
		return -1
	} else if o.Lat > r.Lat {
		return 1
	}

	if o.Long < r.Long {
		return -1
	} else if o.Lat > r.Lat {
		return 1
	}

	return 0
}

type Location struct {
	Lat  float64 `json:"lat" form:"lat"`
	Long float64 `json:"long" form:"long"`

	lat    float64
	long   float64
	cosLat float64
}

// Pre-calculating a latitude and a longitude in radial coordinates.
func (o *Location) PreCalc() *Location {
	o.lat = float64(o.Lat) * RadialCoefficient
	o.long = float64(o.Long) * RadialCoefficient

	o.cosLat = math.Cos(o.lat)

	return o
}

// Calculating a great-circle distance between two places.
// Based on a code from http://www.movable-type.co.uk/scripts/latlong.html
func (o Location) Distance(r Location) float64 {
	a1 := math.Sin((r.lat - o.lat) / 2)
	a2 := math.Sin((r.long - o.long) / 2)

	a := a1*a1 + a2*a2*o.cosLat*r.cosLat

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthMidRadius * c
}

func (o Location) String() string {
	return fmt.Sprint("lat: ", o.Lat, "; long: ", o.Long)
}

func (o Location) Validate() error {
	var errList errCode.ErrList

	if o.Lat < -90. || o.Lat > 90. {
		errList.Add(errors.New("lat: out of bounds"))
	}

	if o.Long < -180. || o.Long > 180. {
		errList.Add(errors.New("long: out of bounds"))
	}

	return errList.Result()
}

func (o Location) Empty() bool {
	return o.Lat == 0 && o.Long == 0
}
