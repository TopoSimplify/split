package split

import (
	"simplex/rng"
	"simplex/lnr"
	"simplex/node"
	"github.com/intdxdt/geom"
	"github.com/intdxdt/rtree"
)

//split hull at vertex with
//maximum_offset offset -- k
func AtScoreSelection(self lnr.Linear, hull *node.Node, gfn geom.GeometryFn) (*node.Node, *node.Node) {
	i, j := hull.Range.I(), hull.Range.J()
	k, _ := self.Score(self, hull.Range)
	// -------------------------------------------
	// i..[ha]..k..[hb]..j
	ha := node.NewFromPolyline(self.Polyline(), rng.NewRange(i, k), gfn)
	hb := node.NewFromPolyline(self.Polyline(), rng.NewRange(k, j), gfn)
	// -------------------------------------------
	return ha, hb
}

//split hull at indexes (index, index, ...)
func AtIndex(self lnr.Linear, hull *node.Node, idxs []int, gfn geom.GeometryFn) []*node.Node {
	//formatter:off
	var pln = self.Polyline()
	var ranges = hull.Range.Split(idxs)
	var subHulls = make([]*node.Node, 0)
	for _, r := range ranges {
		subHulls = append(subHulls, node.NewFromPolyline(pln, r, gfn))
	}
	return subHulls
}

//split hull based on score selected vertex
func SplitNodesInDB(self lnr.Linear, nodedb *rtree.RTree, selections *node.Nodes, gfn geom.GeometryFn) {
	selections.Reverse()
	var que = self.NodeQueue()
	for _, hull := range selections.DataView() {
		var ha, hb = AtScoreSelection(self, hull, gfn)
		nodedb.Remove(hull)

		que.AppendLeft(hb)
		que.AppendLeft(ha)
	}
	//empty selections
	selections.Empty()
}
