package split

import (
	"simplex/rng"
	"simplex/lnr"
	"simplex/node"
	"github.com/intdxdt/geom"
)

//split hull at vertex with
//maximum_offset offset -- k
func AtScoreSelection(self lnr.Linear, hull *node.Node, gfn geom.GeometryFn) (*node.Node, *node.Node) {
	i, j := hull.Range.I(), hull.Range.J()
	k, _ := self.Score(self, hull.Range)
	// -------------------------------------------
	// i..[ha]..k..[hb]..j
	ha := node.New(self.Polyline(), rng.NewRange(i, k), gfn)
	hb := node.New(self.Polyline(), rng.NewRange(k, j), gfn)
	// -------------------------------------------
	return ha, hb
}

//split hull at indexes (index, index, ...)
func AtIndex(self lnr.Linear, hull *node.Node, idxs []int, gfn geom.GeometryFn) []*node.Node {
	//formatter:off
	var pln = self.Polyline()
	var ranges = hull.Range.Split(idxs)
	var sub_hulls = make([]*node.Node, 0)
	for _, r := range ranges {
		sub_hulls = append(sub_hulls, node.New(pln, r, gfn))
	}
	return sub_hulls
}
