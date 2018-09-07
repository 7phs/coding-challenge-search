package model

import (
	"math/rand"
	"strconv"
)

type MockSearchDataSource struct {
	Items ItemsList

	CallCount int
}

func NewMockSearchDataSource(count int) *MockSearchDataSource {
	return (&MockSearchDataSource{}).Generate(count)
}

func (o *MockSearchDataSource) Generate(count int) *MockSearchDataSource {
	for i := 0; i < count; i++ {
		o.Items.Add(&Item{
			Id:   int64(rand.Int()),
			Name: "name - " + strconv.Itoa(rand.Int()),
			Location: Location{
				Lat:  float64(rand.Int31n(180)) - 90.,
				Long: float64(rand.Int31n(360)) - 180.,
			},
		})
	}

	return o
}

func (o *MockSearchDataSource) Load() (ItemsList, error) {
	return o.Items, nil
}

func (o *MockSearchDataSource) List(filter *SearchFilter, paging *Paging) (ItemsList, error) {
	o.CallCount++

	start, limit, err := paging.StartLimit(len(o.Items))
	if err != nil {
		return nil, err
	}

	return o.Items[start:limit], nil
}
