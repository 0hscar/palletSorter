package rectanglesort

import (
	"palletSorter/pkg/types"
	"sort"
)

// type Rectangle struct {
// 	Width  int
// 	Height int
// }

func SortRectangles(rectangles []types.Rectangle) {
	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Height > rectangles[j].Width
	})
}
