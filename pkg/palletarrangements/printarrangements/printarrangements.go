package printarrangements

import (
	"fmt"
	"palletSorter/pkg/types"
)

func PrintArrangementASCII(arrangement []types.PlacedRectangle, K, Q int) {
	// Create a 2D grid
	grid := make([][]string, Q)
	for i := range grid {
		grid[i] = make([]string, K)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	// Fill in rectangles with numbers
	for i, rect := range arrangement {
		symbol := fmt.Sprintf("%d", i+1)
		for y := rect.Y; y < rect.Y+rect.Height; y++ {
			for x := rect.X; x < rect.X+rect.Width; x++ {
				grid[y][x] = symbol
			}
		}
	}

	// Print the grid
	fmt.Println("ASCII Visualization:")
	for i := 0; i < Q; i++ {
		for j := 0; j < K; j++ {
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}
}
