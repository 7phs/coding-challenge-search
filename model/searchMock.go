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
		})
	}

	return o
}

func (o *MockSearchDataSource) List(filter *SearchFilter, paging *Paging) (ItemsList, error) {
	o.CallCount++

	start, limit, err := paging.StartLimit(len(o.Items))
	if err != nil {
		return nil, err
	}

	return o.Items[start:limit], nil
}
