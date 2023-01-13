package main

import (
	"encoding/json"
	"net/http"
)

type flight struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type flightPath struct {
	Path []flight `json:"path"`
}

func calculateFlightPath(w http.ResponseWriter, r *http.Request) {
	var flights []flight
	err := json.NewDecoder(r.Body).Decode(&flights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// sort the flights to find the total flight path
	path := flightPath{}
	path.Path = append(path.Path, flights[0])
	for i := 0; i < len(flights)-1; i++ {
		if flights[i].Destination != flights[i+1].Source {
			path.Path = append(path.Path, flights[i+1])
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(path)
}

func main() {
	http.HandleFunc("/calculate", calculateFlightPath)
	http.ListenAndServe(":8080", nil)
}
