package server

import (
	"fmt"
	"palletSorter/pkg/palletarrangements/find3darrangements"
	"palletSorter/pkg/types"
	"sync"
)

type ViewerData struct {
	Cubes	[]types.PlacedCube
	Width	int
	Height	int
	Depth	int
	mu		sync.RWMutex
}

func (vd *ViewerData) AddCube(cube types.Cube) error {
	vd.mu.Lock()
	defer vd.mu.Unlock()

	tempCubes := make([]types.Cube, len(vd.Cubes))

	for i, placedCube := range vd.Cubes {
		tempCubes[i] = types.Cube{
			Width:  placedCube.Width,
			Height: placedCube.Height,
			Depth:  placedCube.Depth,
		}
	}

	newCubes := append(tempCubes, cube)
	newResult := find3darrangements.Find3DArrangements(newCubes, vd.Width, vd.Height, vd.Depth)

	if len(newResult) == 0 {
		return fmt.Errorf("unable to fit new cube in container")
	}

	vd.Cubes = make([]types.PlacedCube, len(newResult))
	for i, resultCube := range newResult {
		vd.Cubes[i] = resultCube
	}
	return nil
	// vd.Cubes = append(vd.Cubes, cube)
}

func (vd *ViewerData) GetCubes() []types.PlacedCube {
    vd.mu.RLock()
    defer vd.mu.RUnlock()
    cubesCopy := make([]types.PlacedCube, len(vd.Cubes))
    copy(cubesCopy, vd.Cubes)
    return cubesCopy
}
