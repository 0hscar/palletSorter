package savearrangements

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"palletSorter/pkg/types"
)

func SaveArrangementImage(arrangement []types.PlacedRectangle, K, Q int, filename string) error {
	// Create a new image
	img := image.NewRGBA(image.Rect(0, 0, K*50, Q*50)) // Scale by 50 for better visibility

	// Fill background with white
	for y := 0; y < Q*50; y++ {
		for x := 0; x < K*50; x++ {
			img.Set(x, y, color.White)
		}
	}

	// Generate random colors for rectangles
	colors := make([]color.Color, len(arrangement))
	for i := range colors {
		colors[i] = color.RGBA{
			uint8(rand.Intn(200) + 55),
			uint8(rand.Intn(200) + 55),
			uint8(rand.Intn(200) + 55),
			255,
		}
	}

	// Draw rectangles
	for i, rect := range arrangement {
		// Fill rectangle
		for y := rect.Y * 50; y < (rect.Y+rect.Height)*50; y++ {
			for x := rect.X * 50; x < (rect.X+rect.Width)*50; x++ {
				img.Set(x, y, colors[i])
			}
		}

		// Draw border
		for y := rect.Y * 50; y < (rect.Y+rect.Height)*50; y++ {
			img.Set(rect.X*50, y, color.Black)
			img.Set((rect.X+rect.Width)*50-1, y, color.Black)
		}
		for x := rect.X * 50; x < (rect.X+rect.Width)*50; x++ {
			img.Set(x, rect.Y*50, color.Black)
			img.Set(x, (rect.Y+rect.Height)*50-1, color.Black)
		}
	}

	// Save to file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
