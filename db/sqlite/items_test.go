package sqlite

import (
	"database/sql"
	"testing"

	"github.com/7phs/coding-challenge-search/model"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestItems_Load(t *testing.T) {
	connection := "../../test/data/test-data.sqlite3"

	db, err := sql.Open("sqlite3", connection)
	if !assert.Nil(t, err, "failed to connect to %s", connection) {
		return
	}

	items := NewItems(db)

	exist, err := items.Load()
	if !assert.Nil(t, err, "failed to load") {
		return
	}

	expected := model.ItemsList{
		{
			Id:       1,
			Name:     "item1",
			Location: model.Location{Lat: 50, Long: -5.1},
			Url:      "url1",
			Imgs:     []string{"img1.1", "img1.2"},
		},
		{
			Id:       2,
			Name:     "item2",
			Location: model.Location{Lat: 51.5, Long: -2.5},
			Imgs:     []string{"img2.1", "img2.2"},
		},
		{
			Id:       3,
			Location: model.Location{Lat: 52.15, Long: 0},
			Url:      "url3",
			Imgs:     []string{"img3.1", "img3.2"},
		},
		{
			Id:       4,
			Name:     "item4",
			Location: model.Location{Lat: 53.9, Long: 1.12},
			Url:      "url4",
			Imgs:     []string{"img4.1", "img4.2"},
		},
		{
			Id:       5,
			Name:     "item5",
			Location: model.Location{Lat: 54.76, Long: 2.45},
			Url:      "url5",
		},
	}
	assert.Equal(t, expected, exist)
}
