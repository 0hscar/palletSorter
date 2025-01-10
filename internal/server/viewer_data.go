package server

import (
	"fmt"
	"palletSorter/pkg/palletarrangements/find3darrangements"
	"palletSorter/pkg/types"
	"sync"
)

type ViewerData struct {
	Cubes  []types.PlacedCube
	Width  int
	Height int
	Depth  int
	mu     sync.RWMutex
}

func (vd *ViewerData) arrangeCubes(cubes []types.Cube, width, height, depth int) ([]types.PlacedCube, error) {
	result := find3darrangements.Find3DArrangements(cubes, width, height, depth)
	if len(result) == 0 {
		return nil, fmt.Errorf("unable to fit cubes in container")
	}
	return result, nil
}

func (vd *ViewerData) getCubesWithoutPosition() []types.Cube {
	cubes := make([]types.Cube, len(vd.Cubes))
	for i, placedCube := range vd.Cubes {
		cubes[i] = types.Cube{
			Width:  placedCube.Width,
			Height: placedCube.Height,
			Depth:  placedCube.Depth,
		}
	}
	return cubes
}

func (vd *ViewerData) AddCube(cube types.Cube) error {
	vd.mu.Lock()
	defer vd.mu.Unlock()

	currentCubes := vd.getCubesWithoutPosition()
	newCubes := append(currentCubes, cube)

	result, err := vd.arrangeCubes(newCubes, vd.Width, vd.Height, vd.Depth)

	if err != nil {
		return fmt.Errorf("Failed to add cube: %w", err)
	}

	vd.Cubes = result
	return nil
}

func (vd *ViewerData) UpdateContainerSize(width, height, depth int) error {
	vd.mu.Lock()
	defer vd.mu.Unlock()

	// Validate
	if width <= 0 || height <= 0 || depth <= 0 {
		return fmt.Errorf("invalid container dimensions")
	}

	currentCubes := vd.getCubesWithoutPosition()
	result, err := vd.arrangeCubes(currentCubes, width, height, depth)
	if err != nil {
		return fmt.Errorf("failed to resize container: %w", err)
	}

	// new size and arrangements
	vd.Width = width
	vd.Height = height
	vd.Depth = depth
	vd.Cubes = result
	return nil
}

func (vd *ViewerData) GetCubes() []types.PlacedCube {
	vd.mu.RLock()
	defer vd.mu.RUnlock()
	cubesCopy := make([]types.PlacedCube, len(vd.Cubes))
	copy(cubesCopy, vd.Cubes)
	return cubesCopy
}
