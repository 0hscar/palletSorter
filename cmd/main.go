package main

import (
	"fmt"
	"log"
	"palletSorter/internal/server"
	"palletSorter/pkg/palletarrangements/find3darrangements"

	// "palletSorter/pkg/palletarrangements/findarrangements"
	// "palletSorter/pkg/palletarrangements/printarrangements"
	"palletSorter/pkg/palletarrangements/save3darrangements"
	// "palletSorter/pkg/palletarrangements/savearrangements"
	"palletSorter/pkg/types"
)

func main() {
	// Modify how many cubes and their sizes. TODO: Configurable at runtime
	cubes := []types.Cube{
		{Width: 2, Height: 3, Depth: 2},
		{Width: 1, Height: 2, Depth: 3},
		{Width: 2, Height: 2, Depth: 2},
		{Width: 3, Height: 1, Depth: 1},
	}

	// Pallet size. TODO: Configurable at runtime
	width := 6
	height := 4
	depth := 4 // z axis, height from the ground

	// rectangles := []types.Rectangle{
	// 	{Width: 3, Height: 2},
	// 	{Width: 1, Height: 6},
	// 	{Width: 2, Height: 2},
	// 	{Width: 3, Height: 1},
	// 	{Width: 2, Height: 2},
	// }

	// K := 5
	// Q := 6
	// result := findarrangements.FindArrangements(rectangles, K, Q)
	// if result == nil {
	// 	fmt.Println("No arrangement possible")
	// } else {
	// 	fmt.Println("Arrangement found: ", result)
	// }

	// printarrangements.PrintArrangementASCII(result, K, Q)

	// Save graphical visualization
	// err := savearrangements.SaveArrangementImage(result, K, Q, "arrangement.png")
	// if err != nil {
	// 	fmt.Println("Error saving image:", err)
	// 	return
	// }
	// fmt.Println("Image saved as arrangement.png")

	result3d := find3darrangements.Find3DArrangements(cubes, width, height, depth)
	if result3d == nil {
		fmt.Println("No 3D arrangement possible")
		return
	}

	fmt.Println("3D Arrangement found!")
	for i, cube := range result3d {
		fmt.Printf("Cube %d: Position(%d,%d,%d) Size(%d,%d,%d)\n",
			i+1, cube.X, cube.Y, cube.Z, cube.Width, cube.Height, cube.Depth)
	}

	err3d := save3darrangements.Save3DArrangementImage(result3d, width, height, depth, "3d_arrangement.png")
	if err3d != nil {
		fmt.Println("Error saving 3D visualization:", err3d)
		return
	}
	fmt.Println("3D visualization saved as 3d_arrangement.png")

	viewerData := server.ViewerData{
		Cubes:  result3d,
		Width:  width,
		Height: height,
		Depth:  depth,
	}

	if err3d := server.StartViewer(viewerData, ":8080"); err3d != nil {
		log.Fatal("Server error: ", err3d)
	}

}
