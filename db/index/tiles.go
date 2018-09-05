package index

import (
	"os"
	"sync"

	"github.com/7phs/coding-challenge-search/model"
)

const (
	defaultPrecision   = 1000
	defaultDiscret     = 10
	defaultMaxDistance = 20037.5
)

type Tiles struct {
	root      *TileNode
	edgeNodes TileNodeList
}

func NewIndexTiles() *Tiles {
	return &Tiles{
		root: NewTileNode(defaultDiscret * defaultPrecision),
	}
}

func (o *Tiles) Add(record *model.Item) {
	record.Location.PreCalc()

	newEdgeNode := o.root.createEdgeNode(record)
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
	maxDistance := o.edgeNodes.MaxDistance()
	if maxDistance == 0 {
		maxDistance = defaultMaxDistance
	}

	o.edgeNodes.Finish(maxDistance)
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
