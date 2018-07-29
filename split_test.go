package split

import (
	"time"
	"testing"
	"github.com/franela/goblin"
	"github.com/TopoSimplify/dp"
	"github.com/TopoSimplify/opts"
	"github.com/TopoSimplify/offset"
	"github.com/intdxdt/iter"
)

//@formatter:off
func TestSplitHull(t *testing.T) {
	var g = goblin.Goblin(t)
	var id = iter.NewIgen()
	g.Describe("test split hull", func() {
		g.It("should test split", func() {

			g.Timeout(1 * time.Hour)
			options := &opts.Opts{
				Threshold:              50.0,
				MinDist:                20.0,
				RelaxDist:              30.0,
				NonPlanarSelf:          false,
				PlanarSelf:             true,
				AvoidNewSelfIntersects: true,
				GeomRelation:           true,
				DistRelation:           false,
				DirRelation:            false,
			}

			// self.relates = relations(self)
			var wkt = "LINESTRING ( 860 390, 810 360, 770 400, 760 420, 800 440, 810 470, 850 500, 810 530, 780 570, 760 530, 720 530, 710 500, 650 450 )"
			var coords = linearCoords(wkt)
			var n = len(coords) - 1
			var homo = dp.New(coords, options, offset.MaxOffset)
			var hull = createHulls(id, [][]int{{0, n}}, coords)[0]

			ha, hb := AtScoreSelection(id, &hull, homo.Score, hullGeom)

			g.Assert(ha.Range.AsSlice()).Equal([]int{0, 8})
			g.Assert(hb.Range.AsSlice()).Equal([]int{8, len(coords) - 1})

			splits := AtIndex(id, &ha, []int{3, 6}, hullGeom)
			g.Assert(len(splits)).Equal(3)
			g.Assert(splits[0].Range.AsSlice()).Equal([]int{0, 3})
			g.Assert(splits[1].Range.AsSlice()).Equal([]int{3, 6})
			g.Assert(splits[2].Range.AsSlice()).Equal([]int{6, 8})

			splits = AtIndex(id, &hull, []int{ha.Range.I, ha.Range.J, hb.Range.I, hb.Range.J}, hullGeom)

			g.Assert(len(splits)).Equal(2)
			splits = AtIndex(id, &hull, []int{ha.Range.I, ha.Range.J, hb.Range.I, hb.Range.I - 1, hb.Range.J}, hullGeom)
			g.Assert(len(splits)).Equal(3)

			splits = AtIndex(id, &ha, []int{3, 6, 1, 2, 5, 6}, hullGeom)
			g.Assert(len(splits)).Equal(6)
		})
	})
}
