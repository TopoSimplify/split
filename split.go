package split

import (
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/lnr"
	"github.com/TopoSimplify/node"
	"github.com/intdxdt/geom"
)

//split hull at vertex with
//maximum_offset offset -- k
func AtScoreSelection(hull *node.Node, scoreFn lnr.ScoreFn, gfn geom.GeometryFn) (node.Node, node.Node) {
	var coordinates = hull.Coordinates()
	var rg = hull.Range
	var i, j = rg.I, rg.J
	var k, _ = scoreFn(coordinates)
	var idx = rg.I + k
	// ---------------------------------------------------------------
	var idA, idB = hull.SubNodeIds()
	// i..[ha]..k..[hb]..j
	var ha = node.CreateNode(coordinates[0:k+1], rng.Range(i, idx), gfn, idA)
	var hb = node.CreateNode(coordinates[k:], rng.Range(idx, j), gfn,    idB)
	ha.Instance, hb.Instance = hull.Instance, hull.Instance
	// ---------------------------------------------------------------
	return ha, hb
}

//split hull at indices (index, index, ...)
func AtIndex(hull *node.Node, indices []int, gfn geom.GeometryFn) []node.Node {
	//formatter:off
	var coordinates = hull.Coordinates()
	var ranges = hull.Range.Split(indices)
	var subHulls = make([]node.Node, 0, len(ranges))
	var I = hull.Range.I
	for _, r := range ranges {
		subHulls = append(subHulls, node.CreateNode(coordinates[r.I-I:r.J-I+1], r, gfn))
	}
	return subHulls
}
