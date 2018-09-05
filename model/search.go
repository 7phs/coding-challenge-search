package model

type SearchFilter struct {
	Keywords *SearchKeyword
	Location Location
}

type SearchDataSource interface {
	List(*SearchFilter) (ItemsList, error)
}

type Search struct {
	dataSource SearchDataSource
}

func NewSearch(dataSource SearchDataSource) *Search {
	return &Search{
		dataSource: dataSource,
	}
}
