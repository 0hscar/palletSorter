package find3darrangements

import (
	"palletSorter/pkg/types"
)

func Find3DArrangements(cubes []types.Cube, width, height, depth int) []types.PlacedCube {
	doesFit := func(x, y, z, w, h, d int, placedCubes []types.PlacedCube) bool {
		if x+w > width || y+h > height || z+d > depth {
			return false
		}
		for _, cube := range placedCubes {
			if !(x+w <= cube.X || cube.X+cube.Width <= x ||
				y+h <= cube.Y || cube.Y+cube.Height <= y ||
				z+d <= cube.Z || cube.Z+cube.Depth <= z) {
				return false
			}
		}
		return true
	}

	var placeCubes func(index int, placedCubes []types.PlacedCube) []types.PlacedCube
	placeCubes = func(index int, placedCubes []types.PlacedCube) []types.PlacedCube {
		if index == len(cubes) {
			return placedCubes
		}

		w, h, d := cubes[index].Width, cubes[index].Height, cubes[index].Depth

		// positions and orientations
		for z := 0; z <= depth-d; z++ {
			for y := 0; y <= height-h; y++ {
				for x := 0; x <= width-w; x++ {
					if doesFit(x, y, z, w, h, d, placedCubes) {
						newPlaced := append([]types.PlacedCube{}, placedCubes...)
						newPlaced = append(newPlaced, types.PlacedCube{
							X: x, Y: y, Z: z,
							Width: w, Height: h, Depth: d,
						})

						if result := placeCubes(index+1, newPlaced); result != nil {
							return result
						}
					}
				}
			}
		}
		return nil
	}

	return placeCubes(0, []types.PlacedCube{})
}
