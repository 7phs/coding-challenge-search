package sqlite

import (
	"database/sql"

	"github.com/7phs/coding-challenge-search/db/common"
	"github.com/7phs/coding-challenge-search/model"
)

type QueryItemsLoad struct{}

func (o QueryItemsLoad) Query() string {
	return `
	SELECT
		item_name,
		lat,
		lng,
		item_url,
		img_urls
	FROM
		items
	`
}

func (o QueryItemsLoad) Bind() []interface{} {
	return nil
}

func (o QueryItemsLoad) Scan(row *sql.Rows, record *model.Item) error {
	var imgs string

	err := row.Scan(
		&common.NullString{V: &record.Name},
		&common.NullFloat64{V: &record.Location.Lat},
		&common.NullFloat64{V: &record.Location.Long},
		&common.NullString{V: &record.Url},
		&common.NullString{V: &imgs},
	)

	record.Imgs = common.JsonArrayString(imgs).Unmarshal()

	return err
}
