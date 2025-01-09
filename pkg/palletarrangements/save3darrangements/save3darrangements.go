// OLD, SAVES PNG
package save3darrangements

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"palletSorter/pkg/types"
)

func Save3DArrangementImage(arrangement []types.PlacedCube, width, height, depth int, filename string) error {
	// isometric projection
	scale := 50
	// image dimensions based on isometric projection
	imgWidth := (width + depth) * scale
	imgHeight := (width + depth + height) * scale / 2

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Fill background
	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			img.Set(x, y, color.White)
		}
	}

	// Cubes in isometric projection
	for i, cube := range arrangement {
		drawIsometricCube(img, cube, scale, color.RGBA{
			uint8(30 * (i + 1) % 255),
			uint8(50 * (i + 1) % 255),
			uint8(70 * (i + 1) % 255),
			255,
		})
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}

func drawIsometricCube(img *image.RGBA, cube types.PlacedCube, scale int, col color.Color) {
	// Isometric projection matrices
	cos30 := math.Cos(math.Pi / 6)
	sin30 := math.Sin(math.Pi / 6)

	// 3D --> 2D
	project := func(x, y, z float64) (int, int) {
		px := (x - z) * cos30
		py := y + (x+z)*sin30
		return int(px*float64(scale)) + img.Bounds().Max.X/2,
			int(py*float64(scale)) + img.Bounds().Max.Y/4
	}

	// Calculate vertices for the cube
	vertices := [][3]float64{
		{float64(cube.X), float64(cube.Y), float64(cube.Z)},                                         // 0: front bottom left
		{float64(cube.X + cube.Width), float64(cube.Y), float64(cube.Z)},                            // 1: front bottom right
		{float64(cube.X + cube.Width), float64(cube.Y + cube.Height), float64(cube.Z)},              // 2: front top right
		{float64(cube.X), float64(cube.Y + cube.Height), float64(cube.Z)},                           // 3: front top left
		{float64(cube.X), float64(cube.Y), float64(cube.Z + cube.Depth)},                            // 4: back bottom left
		{float64(cube.X + cube.Width), float64(cube.Y), float64(cube.Z + cube.Depth)},               // 5: back bottom right
		{float64(cube.X + cube.Width), float64(cube.Y + cube.Height), float64(cube.Z + cube.Depth)}, // 6: back top right
		{float64(cube.X), float64(cube.Y + cube.Height), float64(cube.Z + cube.Depth)},              // 7: back top left
	}

	// Project vertices to 2D
	points := make([][2]int, 8)
	for i, v := range vertices {
		points[i][0], points[i][1] = project(v[0], v[1], v[2])
	}

	// Define faces (each face is defined by 4 vertex indices)
	faces := [][4]int{
		{0, 1, 2, 3}, // Front
		{4, 5, 6, 7}, // Back
		{0, 4, 7, 3}, // Left
		{1, 5, 6, 2}, // Right
		{3, 2, 6, 7}, // Top
		{0, 1, 5, 4}, // Bottom
	}

	// Colors for different faces
	faceColors := []color.Color{
		color.RGBA{R: uint8(col.(color.RGBA).R), G: uint8(col.(color.RGBA).G), B: uint8(col.(color.RGBA).B), A: 255},                                              // Front
		color.RGBA{R: uint8(float64(col.(color.RGBA).R) * 0.6), G: uint8(float64(col.(color.RGBA).G) * 0.6), B: uint8(float64(col.(color.RGBA).B) * 0.6), A: 255}, // Back
		color.RGBA{R: uint8(float64(col.(color.RGBA).R) * 0.8), G: uint8(float64(col.(color.RGBA).G) * 0.8), B: uint8(float64(col.(color.RGBA).B) * 0.8), A: 255}, // Left
		color.RGBA{R: uint8(float64(col.(color.RGBA).R) * 0.7), G: uint8(float64(col.(color.RGBA).G) * 0.7), B: uint8(float64(col.(color.RGBA).B) * 0.7), A: 255}, // Right
		color.RGBA{R: uint8(float64(col.(color.RGBA).R) * 0.9), G: uint8(float64(col.(color.RGBA).G) * 0.9), B: uint8(float64(col.(color.RGBA).B) * 0.9), A: 255}, // Top
		color.RGBA{R: uint8(float64(col.(color.RGBA).R) * 0.5), G: uint8(float64(col.(color.RGBA).G) * 0.5), B: uint8(float64(col.(color.RGBA).B) * 0.5), A: 255}, // Bottom
	}

	for i, face := range faces {
		drawFace(img, points, face, faceColors[i])
	}

	edges := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0}, // Front face
		{4, 5}, {5, 6}, {6, 7}, {7, 4}, // Back face
		{0, 4}, {1, 5}, {2, 6}, {3, 7}, // Connecting edges
	}

	for _, edge := range edges {
		drawLine(img,
			points[edge[0]][0], points[edge[0]][1],
			points[edge[1]][0], points[edge[1]][1],
			color.Black)
	}
}

func drawFace(img *image.RGBA, points [][2]int, face [4]int, col color.Color) {
	minX, maxX := img.Bounds().Max.X, img.Bounds().Min.X
	minY, maxY := img.Bounds().Max.Y, img.Bounds().Min.Y

	for _, idx := range face {
		x, y := points[idx][0], points[idx][1]
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if pointInPolygon(x, y, points, face) {
				img.Set(x, y, col)
			}
		}
	}
}

func drawLine(img *image.RGBA, x0, y0, x1, y1 int, col color.Color) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	steep := dy > dx

	if steep {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
	}
	if x0 > x1 {
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx = x1 - x0
	dy = abs(y1 - y0)
	err := dx / 2
	ystep := -1
	if y0 < y1 {
		ystep = 1
	}

	for x := x0; x <= x1; x++ {
		if steep {
			img.Set(y0, x, col)
		} else {
			img.Set(x, y0, col)
		}
		err -= dy
		if err < 0 {
			y0 += ystep
			err += dx
		}
	}
}

func pointInPolygon(x, y int, points [][2]int, face [4]int) bool {
	inside := false
	j := face[len(face)-1]

	for _, i := range face {
		if ((points[i][1] > y) != (points[j][1] > y)) &&
			(x < (points[j][0]-points[i][0])*(y-points[i][1])/(points[j][1]-points[i][1])+points[i][0]) {
			inside = !inside
		}
		j = i
	}

	return inside
}

// function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
