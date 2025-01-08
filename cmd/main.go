package main

import (
	"fmt"
	"palletSorter/pkg/palletarrangements/findarrangements"
	"palletSorter/pkg/palletarrangements/printarrangements"
	"palletSorter/pkg/palletarrangements/savearrangements"
	"palletSorter/pkg/types"
)

func main() {

	rectangles := []types.Rectangle{
		{Width: 3, Height: 2},
		{Width: 1, Height: 6},
		{Width: 2, Height: 2},
		{Width: 3, Height: 1},
		{Width: 2, Height: 2},
	}

	K := 5
	Q := 6
	result := findarrangements.FindArrangements(rectangles, K, Q)
	if result == nil {
		fmt.Println("No arrangement possible")
	} else {
		fmt.Println("Arrangement found: ", result)
	}

	printarrangements.PrintArrangementASCII(result, K, Q)

	// Save graphical visualization
	err := savearrangements.SaveArrangementImage(result, K, Q, "arrangement.png")
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}
	fmt.Println("Image saved as arrangement.png")

}
