package split

import (
	"simplex/rng"
	"simplex/lnr"
	"simplex/node"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
	"github.com/intdxdt/deque"
)

//split hull at vertex with
//maximum_offset offset -- k
func AtScoreSelection(hull *node.Node, scoreFn lnr.ScoreFn, gfn geom.GeometryFn) (*node.Node, *node.Node) {
	var coordinates = hull.Coordinates()
	var rg = hull.Range
	var i, j = rg.I(), rg.J()
	var k, _ = scoreFn(coordinates)
	var rk = rg.Index(k)
	// -------------------------------------------
	// i..[ha]..k..[hb]..j
	ha := node.New(coordinates[0:k+1], rng.NewRange(i, rk), gfn)
	hb := node.New(coordinates[k:], rng.NewRange(rk, j), gfn)
	// -------------------------------------------
	return ha, hb
}

//split hull at indexes (index, index, ...)
func AtIndex(hull *node.Node, idxs []int, gfn geom.GeometryFn) []*node.Node {
	//formatter:off
	var coordinates = hull.Coordinates()
	var ranges = hull.Range.Split(idxs)
	var subHulls = make([]*node.Node, 0)
	var I = hull.Range.I()
	for _, r := range ranges {
		var i, j = r.I()-I, r.J()-I
		var coords = coordinates[i:j+1]
		subHulls = append(subHulls, node.New(coords, r, gfn))
	}
	return subHulls
}

//split hull based on score selected vertex
func SplitNodesInDB(
	que *deque.Deque,
	scoreFn lnr.ScoreFn,
	nodeDB *rtree.RTree,
	selections *node.Nodes,
	gFn geom.GeometryFn,
) {
	selections.Reverse()
	for _, hull := range selections.DataView() {
		var ha, hb = AtScoreSelection(hull, scoreFn, gFn)
		nodeDB.Remove(hull)

		que.AppendLeft(hb)
		que.AppendLeft(ha)
	}
	//empty selections
	selections.Empty()
}
