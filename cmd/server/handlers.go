package main

import (
	"encoding/json"
	"net/http"
)

var storage = NewMemoryStore()

func APIRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	root := map[string]interface{}{
		"kind":       "APIResourceList",
		"apiVersion": "v1",
		"resources": []map[string]string{
			{"name": "gadgets", "kind": "Gadget", "verbs": "get,create,list"},
		},
	}
	json.NewEncoder(w).Encode(root)
}

func ListOrCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(storage.List())

	case "POST":
		var g Gadget
		json.NewDecoder(r.Body).Decode(&g)
		storage.Create(g)
		json.NewEncoder(w).Encode(g)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
