package model

type Item struct {
	Name     string
	Location Location
	Url      string
	Imgs     []string
}

type ItemsList []*Item

func (o *ItemsList) Add(item *Item) {
	*o = append(*o, item)
}
