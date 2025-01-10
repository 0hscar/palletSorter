package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"palletSorter/pkg/types"
	"path/filepath"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("web", "templates", "viewer.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Template error: %v", err),
			http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
		return
	}
}

func handleGetCubes(data *ViewerData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

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
		}
	}
}

func handleAddCube(data *ViewerData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var newCube types.Cube
		if err := json.NewDecoder(r.Body).Decode(&newCube); err != nil { // Check what does &newCube entail
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
			return
		}

		if err := data.AddCube(newCube); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}
