package findarrangements

import (
	"palletSorter/pkg/functions/rectanglesort"
	"palletSorter/pkg/types"
)

func FindArrangements(rectangles []types.Rectangle, K, Q int) []types.PlacedRectangle {
	// sort.Slice(rectangles, func(i, j int) bool {
	// 	return rectangles[i].Height > rectangles[j].Width
	// })

	rectanglesort.SortRectangles(rectangles)

	doesFit := func(x, y, w, h int, placedRects []types.PlacedRectangle) bool {
		if x+w > K || y+h > Q {
			return false
		}
		for _, rect := range placedRects {
			if !(x+w <= rect.X || rect.X+rect.Width <= x ||
				y+h <= rect.Y || rect.Y+rect.Height <= y) {
				return false
			}
		}
		return true
	}

	var placeRectangles func(index int, placedRects []types.PlacedRectangle) []types.PlacedRectangle
	placeRectangles = func(index int, placedRects []types.PlacedRectangle) []types.PlacedRectangle {
		if index == len(rectangles) {
			return placedRects
		}

		w, h := rectangles[index].Width, rectangles[index].Height
		for y := 0; y <= Q-h; y++ {
			for x := 0; x <= K-w; x++ {
				if doesFit(x, y, w, h, placedRects) {
					newPlaced := append([]types.PlacedRectangle{}, placedRects...)
					newPlaced = append(newPlaced, types.PlacedRectangle{X: x, Y: y, Width: w, Height: h})

					if result := placeRectangles(index+1, newPlaced); result != nil {
						return result
					}
				}
			}
		}
		return nil
	}
	result := placeRectangles(0, []types.PlacedRectangle{})
	return result

}
