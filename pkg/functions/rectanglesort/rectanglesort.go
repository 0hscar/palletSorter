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
		if rectangles[i].Height == rectangles[j].Height {
			return rectangles[i].Width > rectangles[j].Width
		}
		return rectangles[i].Height > rectangles[j].Height
	})
}
