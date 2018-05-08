package split

import (
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/intdxdt/geom"
)

//hull geom
func hullGeom(coords []*geom.Point) geom.Geometry {
	var g geom.Geometry

	if len(coords) > 2 {
		g = geom.NewPolygon(coords)
	} else if len(coords) == 2 {
		g = geom.NewLineString(coords)
	} else {
		g = coords[0].Clone()
	}
	return g
}

func linearCoords(wkt string) []*geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createHulls(indxs [][]int, coords []*geom.Point) []*node.Node {
	poly := pln.New(coords)
	hulls := make([]*node.Node, 0)
	for _, o := range indxs {
		var r = rng.NewRange(o[0], o[1])
		var n = node.New(poly.SubCoordinates(r), r, hullGeom)
		hulls = append(hulls, n)
	}
	return hulls
}
