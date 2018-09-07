package index

import (
	"container/heap"
	"sort"
	"sync"

	"github.com/7phs/coding-challenge-search/model"
)

const (
	posibleBorderedNode = 4
)

type TileNode struct {
	level     int
	precision int
	location  model.LocationInt64
	children  TileNodeList

	minDistance float64
	maxDistance float64
	items       *ItemResult
}

func NewTileNode(level int, location model.LocationInt64) *TileNode {
	var (
		precision = defaultPrecision / level
		items     *ItemResult
	)
	if precision < 1 {
		precision = 1
	}

	if level == 1 {
		items = &ItemResult{}
	}

	return &TileNode{
		level:       level,
		precision:   precision,
		location:    location,
		minDistance: defaultMaxDistance,
		items:       items,
	}
}

func (o *TileNode) Depth() int {
	if len(o.children) == 0 {
		return 0
	}

	return 1 + o.children[0].Depth()
}

func (o *TileNode) NodesCount() int {
	if len(o.children) == 0 {
		return 0
	}

	count := 0

	for _, child := range o.children {
		count += 1 + child.NodesCount()
	}

	return count
}

func (o *TileNode) addEdgeNode(location model.Location) *TileNode {
	loc64 := model.NewLocationInt64(location, o.precision)

	child := o.children.Has(loc64)
	if child == nil {
		child = NewTileNode(o.level/defaultDiscret, *loc64.PreCalc())

		o.children.Add(child)
		o.children.Sort()

		if child.level == 1 {
			return child
		}
	} else if child.level == 1 {
		return nil
	}

	return child.addEdgeNode(location)
}

func (o *TileNode) Add(record *model.Item) {
	distance := o.location.Distance(record.Location)

	o.items.Add(record, distance)

	if o.minDistance > distance {
		o.minDistance = distance
	}

	if o.maxDistance < distance {
		o.maxDistance = distance
	}
}

func (o *TileNode) Finish(start, limit float64) {
	o.items.Normalize(start, limit)

	o.items.Sort()
}

func (o *TileNode) Search(location model.Location, distance float64) *TileNodeResult {
	if o.level == 1 {
		return &TileNodeResult{
			items:    o.items,
			distance: distance,
		}
	}

	if len(o.children) == 0 {
		return nil
	}

	var (
		nearestChildren = make(TileNodeHeap, 0, len(o.children))
		result          = make(chan *TileNodeResult)
		nearest         = &TileNodeResult{
			distance: 2 * defaultMaxDistance,
		}
		wait       sync.WaitGroup
		waitResult sync.WaitGroup
	)

	for _, child := range o.children {
		heap.Push(&nearestChildren, NewTileNodeCalculated(child, child.location.Distance(location)))
	}

	for i := 0; i < posibleBorderedNode && i < nearestChildren.Len(); i++ {
		wait.Add(1)
		go func(child *TileNodeCalculated) {
			defer wait.Done()

			result <- child.TileNode.Search(location, child.distance)
		}(heap.Pop(&nearestChildren).(*TileNodeCalculated))
	}

	waitResult.Add(1)
	go func() {
		defer waitResult.Done()

		for res := range result {
			if res.items.Empty() {
				continue
			}

			if res.distance < nearest.distance {
				nearest = res
			}
		}
	}()

	wait.Wait()

	close(result)
	waitResult.Wait()

	return nearest
}

type TileNodeResult struct {
	items    *ItemResult
	distance float64
}

type TileNodeList []*TileNode

func (o *TileNodeList) Add(node *TileNode) {
	*o = append(*o, node)
}

func (o TileNodeList) Factors() (float64, float64) {
	max := 0.
	min := 0.

	if len(o) > 0 {
		min = defaultMaxDistance

		for _, node := range o {
			if min > node.minDistance {
				min = node.minDistance
			}
			if max < node.maxDistance {
				max = node.maxDistance
			}
		}
	}

	if max <= min {
		max = min + 1.
	}

	return min, max - min
}

func (o TileNodeList) Finish() {
	var wait sync.WaitGroup

	start, limit := o.Factors()

	for _, node := range o {
		wait.Add(1)
		go func(node *TileNode) {
			defer wait.Done()

			node.Finish(start, limit)
		}(node)
	}

	wait.Wait()
}

func (o *TileNodeList) Sort() {
	sort.Slice(*o, func(i, j int) bool {
		return (*o)[i].location.Compare(&(*o)[j].location) < 0
	})
}

func (o TileNodeList) Has(r *model.LocationInt64) *TileNode {
	index := sort.Search(len(o), func(i int) bool {
		return o[i].location.Compare(r) >= 0
	})

	if index < len(o) && o[index].location.Compare(r) == 0 {
		return o[index]
	}

	return nil
}

type TileNodeCalculated struct {
	*TileNode

	distance float64
}

func NewTileNodeCalculated(node *TileNode, distance float64) *TileNodeCalculated {
	return &TileNodeCalculated{
		TileNode: node,
		distance: distance,
	}
}

type TileNodeHeap []*TileNodeCalculated

func (h TileNodeHeap) Len() int           { return len(h) }
func (h TileNodeHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h TileNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TileNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*TileNodeCalculated))
}

func (h *TileNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
