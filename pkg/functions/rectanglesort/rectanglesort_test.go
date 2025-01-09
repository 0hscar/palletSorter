package rectanglesort

import (
	"fmt"
	"math/rand"
	"palletSorter/pkg/types"
	"reflect"
	"testing"
)

func BenchmarkSortRectangles(b *testing.B) {
	// rectangles := []types.Rectangle{
	// 	{Width: 2, Height: 4},
	// 	{Width: 5, Height: 3},
	// 	{Width: 6, Height: 6},
	// 	{Width: 3, Height: 8},
	// 	{Width: 1, Height: 2},
	// 	{Width: 7, Height: 10},
	// }

	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			rectangles := make([]types.Rectangle, size)
			for i := 0; i < size; i++ {
				rectangles[i] = types.Rectangle{
					Width:  rand.Intn(100),
					Height: rand.Intn(100),
				}
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				testSlice := make([]types.Rectangle, len(rectangles))
				copy(testSlice, rectangles)
				SortRectangles(testSlice)
			}
		})
	}
}

func TestSortRectangles(t *testing.T) {

	tests := []struct {
		name     string
		input    []types.Rectangle
		expected []types.Rectangle
	}{
		{
			name: "Basic sorting",
			input: []types.Rectangle{
				{Width: 2, Height: 3},
				{Width: 5, Height: 6},
				{Width: 1, Height: 4},
			},
			expected: []types.Rectangle{
				{Width: 5, Height: 6},
				{Width: 1, Height: 4},
				{Width: 2, Height: 3},
			},
		},
		{
			name:     "Empty slice",
			input:    []types.Rectangle{},
			expected: []types.Rectangle{},
		},
		{
			name: "Single Rectangle",
			input: []types.Rectangle{
				{Width: 2, Height: 3},
			},
			expected: []types.Rectangle{
				{Width: 2, Height: 3},
			},
		},
		{
			name: "Equal height",
			input: []types.Rectangle{
				{Width: 5, Height: 6},
				{Width: 3, Height: 6},
				{Width: 4, Height: 6},
			},
			expected: []types.Rectangle{
				{Width: 5, Height: 6},
				{Width: 4, Height: 6},
				{Width: 3, Height: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testInput := make([]types.Rectangle, len(tt.input))
			copy(testInput, tt.input)

			SortRectangles(testInput)

			if !reflect.DeepEqual(testInput, tt.expected) {
				t.Errorf("SortRectangles() = %v, want %v", testInput, tt.expected)
			}

		})
	}

}
