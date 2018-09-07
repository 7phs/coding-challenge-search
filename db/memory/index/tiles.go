package index

import (
	"os"
	"sync"

	"github.com/7phs/coding-challenge-search/model"
)

const (
	defaultPrecision   = 1000000
	defaultDiscret     = 100
	defaultMaxDistance = 20037.5
)

type Tiles struct {
	root      *TileNode
	edgeNodes TileNodeList
}

func NewIndexTiles() *Tiles {
	return &Tiles{
		root: NewTileNode(defaultDiscret*defaultPrecision, model.LocationInt64{}),
	}
}

func (o *Tiles) Add(record *model.Item) {
	newEdgeNode := o.root.addEdgeNode(record.Location)
	if newEdgeNode != nil {
		o.edgeNodes.Add(newEdgeNode)
	}

	var wait sync.WaitGroup
	for _, edgeNode := range o.edgeNodes {
		wait.Add(1)
		go func(edgeNode *TileNode) {
			defer wait.Done()

			edgeNode.Add(record)
		}(edgeNode)
	}
	wait.Wait()
}

func (o *Tiles) Finish() {
	o.edgeNodes.Finish()
}

func (o *Tiles) Search(filter *model.SearchFilter) (Result, error) {
	if filter.Location.Empty() {
		return nil, os.ErrInvalid
	}

	filter.Location.PreCalc()

	result := o.root.Search(filter.Location, 0.)
	if result.items == nil {
		return nil, os.ErrNotExist
	}

	return result.items, nil
}
