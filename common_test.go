package split

import (
	"github.com/TopoSimplify/pln"
	"github.com/TopoSimplify/rng"
	"github.com/TopoSimplify/node"
	"github.com/intdxdt/geom"
)

//hull geom
func hullGeom(coordinates []geom.Point) geom.Geometry {
	var g geom.Geometry
	if len(coordinates) > 2 {
		g = geom.NewPolygon(coordinates)
	} else if len(coordinates) == 2 {
		g = geom.NewLineString(coordinates)
	} else {
		g = coordinates[0]
	}
	return g
}

func linearCoords(wkt string) []geom.Point {
	return geom.NewLineStringFromWKT(wkt).Coordinates()
}

func createHulls(indxs [][]int, coords []geom.Point) []node.Node {
	poly := pln.New(coords)
	hulls := make([]node.Node, 0)
	for _, o := range indxs {
		var r = rng.Range(o[0], o[1])
		var n = node.CreateNode(poly.SubCoordinates(r), r, hullGeom)
		hulls = append(hulls, n)
	}
	return hulls
}
