package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"mime"
	"net/http"
	"palletSorter/pkg/palletarrangements/find3darrangements"
	"palletSorter/pkg/types"
	"path/filepath"
	"sync"
)

type ViewerData struct {
	Cubes  []types.PlacedCube
	Width  int
	Height int
	Depth  int
	mu     sync.RWMutex // Add mutex for thread-safe operations
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

func init() {
	// Register the correct MIME type for JavaScript files
	mime.AddExtensionType(".js", "application/javascript")
}

func StartViewer(data *ViewerData, port string) error {
	// Debug: Print working directory
	absPath, _ := filepath.Abs(".")
	fmt.Printf("Working directory: %s\n", absPath)

	// Serve static files with proper MIME types
	fileServer := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set correct content type for .js files
		if filepath.Ext(r.URL.Path) == ".js" {
			w.Header().Set("Content-Type", "application/javascript")
		}
		http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
	}))

	// Serve the main viewer page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplPath := filepath.Join("web", "templates", "viewer.html")
		absPath, _ := filepath.Abs(tmplPath)
		fmt.Printf("Looking for template at: %s\n", absPath)

		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Template error: %v", err), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
			return
		}
	})

	// API endpoint to serve the cube data
	http.HandleFunc("/api/cubes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			response := struct {
				Cubes  []types.PlacedCube `json:"cubes"`
				Width  int                `json:"width"`
				Height int                `json:"height"`
				Depth  int                `json:"depth"`
			}{
				Cubes:  data.GetCubes(),
				Width:  data.Width,
				Height: data.Height,
				Depth:  data.Depth,
			}
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, fmt.Sprintf("JSON encoding error: %v", err), http.StatusInternalServerError)
				return
			}
		}
	})

	http.HandleFunc("/api/cubes/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var newCube types.Cube
		if err := json.NewDecoder(r.Body).Decode(&newCube); err != nil {
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
			return
		}

		if err := data.AddCube(newCube); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusCreated)

	})

	// Print server start message
	fmt.Printf("3D Viewer server starting at http://localhost%s\n", port)
	fmt.Println("Use Ctrl+C to stop the server")

	// Start the server
	return http.ListenAndServe(port, nil)
}
