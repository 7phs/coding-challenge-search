package model

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemsList_Has(t *testing.T) {
	var (
		list   ItemsList
		idList []int64
	)

	for i := 0; i < 5; i++ {
		id := rand.Int63n(100)
		idList = append(idList, id)

		list.Add(&Item{
			Id:   id,
			Name: "name " + strconv.Itoa(i),
		})
	}

	list.Sort()

	assert.Len(t, list, len(idList))

	sort.Slice(idList, func(i, j int) bool {
		return idList[i] < idList[j]
	})

	for i, rec := range list {
		assert.Equal(t, idList[i], rec.Id, "%d: sorted id", i+1)
	}

	assert.Equal(t, true, list.Has(&Item{
		Id: idList[rand.Int31n(int32(len(idList)))],
	}))
}

func TestItemWithRateList_Sort(t *testing.T) {
	expected := ItemsList{
		{Id: 123, Name: "name 123"},
		{Id: 123, Name: "name 123"},
		{Id: 125, Name: "name 125"},
		{Id: 127, Name: "name 127"},
	}
	rates := []float64{3.6, 3.6, 12.05, 130.}
	expectedByRate := ItemsList{
		expected[3], expected[2], expected[1], expected[0],
	}

	var exist ItemWithRateList
	for _, idx := range []int{2, 3, 1, 0} {
		exist.Add(
			NewItemWithRate(expected[idx], rates[idx]))
	}

	assert.Equal(t, expected, exist.SortById().ItemsList(), "sort by id")
	assert.Equal(t, expectedByRate, exist.Sort().ItemsList(), "sort by rate")

	copied := exist.Copy()
	copied.SortById()
	assert.Equal(t, expectedByRate, exist.ItemsList(), "stay not changing if sorting copied")
	assert.Equal(t, expected, copied.ItemsList(), "copied sort by id")

}

func TestItemWithRateList_Intersect(t *testing.T) {
	testSuites := []*struct {
		in1 ItemsList
		in2 ItemsList
		exp ItemsList
	}{
		{
			in1: ItemsList{
				{Id: 123, Name: "name 123"},
				{Id: 125, Name: "name 125"},
				{Id: 127, Name: "name 127"},
				{Id: 128, Name: "name 128"},
			},
			in2: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
			exp: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 128, Name: "name 128"},
			},
		},
		{
			in2: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
		},
		{
			in1: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
		},
	}

	convert := func(itemsList ItemsList) (result ItemWithRateList) {
		if len(itemsList) == 0 {
			return nil
		}

		for _, rec := range itemsList {
			result.Add(NewItemWithRate(rec, rand.Float64()))
		}
		return
	}

	for i, test := range testSuites {
		exist := convert(test.in1).
			Intersect(convert(test.in2)).
			ItemsList()

		assert.Equal(t, test.exp, exist, "%d: intersect", i+1)
	}

}

func TestItemWithRateList_Join(t *testing.T) {
	testSuites := []*struct {
		in1 ItemsList
		in2 ItemsList
		exp ItemsList
	}{
		{
			in1: ItemsList{
				{Id: 123, Name: "name 123"},
				{Id: 125, Name: "name 125"},
				{Id: 127, Name: "name 127"},
				{Id: 128, Name: "name 128"},
			},
			in2: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
			exp: ItemsList{
				{Id: 123, Name: "name 123"},
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 127, Name: "name 127"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
		},
		{
			in2: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
			exp: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
		},
		{
			in1: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
			exp: ItemsList{
				{Id: 125, Name: "name 125"},
				{Id: 126, Name: "name 126"},
				{Id: 128, Name: "name 128"},
				{Id: 129, Name: "name 129"},
			},
		},
	}

	convert := func(itemsList ItemsList) (result ItemWithRateList) {
		if len(itemsList) == 0 {
			return nil
		}

		for _, rec := range itemsList {
			result.Add(NewItemWithRate(rec, rand.Float64()))
		}
		return
	}

	for i, test := range testSuites {
		exist := convert(test.in1).
			Join(convert(test.in2)).
			ItemsList()

		assert.Equal(t, test.exp, exist, "%d: join", i+1)
	}

}
