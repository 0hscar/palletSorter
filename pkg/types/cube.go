package types

type Cube struct {
	Width  int
	Height int
	Depth  int
}

type PlacedCube struct {
	X      int `json:"X"`
	Y      int `json:"Y"`
	Z      int `json:"Z"`
	Width  int `json:"Width"`
	Height int `json:"Height"`
	Depth  int `json:"Depth"`
	// Color  string `json:"Color"`
}
