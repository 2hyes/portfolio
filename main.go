package main

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	// CORS 헤더 설정
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	projects := []Project{
		{"Project A", "Description of Project A"},
		{"Project B", "Description of Project B"},
		{"Project C", "Description of Project C"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func main() {
	http.HandleFunc("/api/projects", projectsHandler)
	http.ListenAndServe(":8080", nil)
}
