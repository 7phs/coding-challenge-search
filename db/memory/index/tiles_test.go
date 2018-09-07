package index

import (
	"testing"

	"github.com/7phs/coding-challenge-search/model"
	"github.com/stretchr/testify/assert"
)

func testLevelsCalc(level, precision int) int {
	depth := 0

	for ; level > 1; depth++ {
		level /= precision
	}

	return depth
}

func testTileData() model.ItemsList {
	data := model.ItemsList{
		{
			Id:       1,
			Location: model.Location{Lat: 55.820423, Long: 49.094051},
		},
		{
			Id:       2,
			Location: model.Location{Lat: 56.109201, Long: 47.234076},
		},
		{
			Id:       3,
			Location: model.Location{Lat: 55.401387, Long: 49.550818},
		},
	}

	for _, rec := range data {
		rec.Location.PreCalc()
	}

	return data
}

func TestTileNode_Add(t *testing.T) {
	location := model.NewLocationInt64(model.Location{Lat: 55.751631, Long: 48.752316}, defaultPrecision).PreCalc()

	tileNode := NewTileNode(1, *location)

	data := testTileData()

	for _, rec := range data {
		tileNode.Add(rec)
	}

	assert.Len(t, tileNode.items.recordsById, len(data))

	expDist := 22.694370
	assert.InEpsilon(t, tileNode.minDistance, expDist, 1e-5)

	expDist = 102.588459
	assert.InEpsilon(t, tileNode.maxDistance, expDist, 1e-5)

	tileNode.Finish(tileNode.minDistance, tileNode.maxDistance-tileNode.minDistance)

	assert.InEpsilon(t, 1.5, tileNode.items.recordsByRate[0].Rate, 1e-5)
	assert.InEpsilon(t, 0.988882, tileNode.items.recordsByRate[1].Rate, 1e-5)
	assert.InEpsilon(t, 0.5, tileNode.items.recordsByRate[2].Rate, 1e-5)
}

func TestTileNode_addEdgeNodeDepth(t *testing.T) {
	location := model.Location{Lat: 55.751631, Long: 48.752316}
	location64 := model.NewLocationInt64(location, defaultPrecision).PreCalc()

	for level := defaultDiscret * defaultPrecision; level > 1; level /= defaultDiscret {
		tileNode := NewTileNode(level, *location64)

		tileNode.addEdgeNode(location)

		expected := testLevelsCalc(level, defaultDiscret)
		assert.Equal(t, expected, tileNode.Depth(), "%d: depth", level)
		assert.Equal(t, expected, tileNode.NodesCount(), "%d: nodes count", level)
	}
}

func TestTileNode_rootNode(t *testing.T) {
	startLevel := defaultDiscret * defaultPrecision
	root := NewTileNode(startLevel, model.LocationInt64{})

	location1 := model.Location{Lat: 55.751631, Long: 48.752316}
	edgeNode := root.addEdgeNode(location1)

	location2 := model.Location{Lat: 56.751631, Long: 49.752316}
	root.addEdgeNode(location2)

	expected := testLevelsCalc(startLevel, defaultDiscret)
	assert.Equal(t, expected, root.Depth(), "depth")
	assert.Equal(t, expected*2, root.NodesCount(), "nodes count")
	assert.Equal(t, 557516, edgeNode.location.Lat, "edge node lat")

	location1.Lat += 6e-6
	edgeNode = root.addEdgeNode(location1)

	assert.Equal(t, expected, root.Depth(), "depth")
	assert.Equal(t, expected*2, root.NodesCount(), "nodes count")
	assert.Nil(t, edgeNode, "already exist")

	location1.Lat += 6e-4
	edgeNode = root.addEdgeNode(location1)

	assert.Equal(t, expected, root.Depth(), "depth")
	assert.Equal(t, expected*2+1, root.NodesCount(), "nodes count")
	assert.Equal(t, 557522, edgeNode.location.Lat, "edge node lat")

	location1.Lat += 6e-2
	edgeNode = root.addEdgeNode(location1)

	assert.Equal(t, expected, root.Depth(), "depth")
	assert.Equal(t, expected*2+3, root.NodesCount(), "nodes count")
	assert.Equal(t, 558122, edgeNode.location.Lat, "edge node lat")
}

func TestTileNode_Search(t *testing.T) {
	startLevel := defaultDiscret * defaultPrecision
	root := NewTileNode(startLevel, model.LocationInt64{})

	data := testTileData()

	var edgeNodes TileNodeList

	for _, rec := range data {
		if edgeNode := root.addEdgeNode(rec.Location); edgeNode != nil {
			edgeNodes.Add(edgeNode)
		}
	}

	for _, rec := range data {
		for _, node := range edgeNodes {
			node.Add(rec)
		}
	}

	start, limit := edgeNodes.Factors()
	assert.InEpsilon(t, 0.004085, start, 1e-4, "start - %f", start)
	assert.InEpsilon(t, 164.940483, limit, 1e-5, "limit - %f", limit)

	edgeNodes.Finish()

	assert.Len(t, edgeNodes, 3)

	for _, node := range edgeNodes {
		assert.Len(t, node.items.recordsById, 3)
	}

	for i, node := range edgeNodes {
		assert.Equal(t, node.items.recordsByRate[0].Item, data[i])
	}

	searchLocation := data[0].Location
	searchLocation.Lat += 6e-6
	exist := root.Search(searchLocation, 10)

	assert.Equal(t,
		edgeNodes[0].items.recordsByRate.ItemsList(),
		exist.items.Items(&model.Paging{Limit: 20}))
}

func TestNewIndexTiles(t *testing.T) {
	idx := NewIndexTiles()

	data := testTileData()

	for _, rec := range data {
		idx.Add(rec)
	}

	idx.Finish()

	_, err := idx.Search(&model.SearchFilter{})
	assert.NotNil(t, err)

	searchLocation := data[1].Location
	searchLocation.Lat += 6e-6
	exist, err := idx.Search(&model.SearchFilter{
		Location: searchLocation,
	})

	assert.Nil(t, err)

	assert.Equal(t,
		idx.edgeNodes[1].items.recordsByRate.ItemsList(),
		exist.Items(&model.Paging{Limit: 20}))
}
