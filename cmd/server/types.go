package main

type Gadget struct {
	Kind       string                 `json:"kind"`
	APIVersion string                 `json:"apiVersion"`
	Metadata   Metadata               `json:"metadata"`
	Spec       map[string]interface{} `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type GadgetList struct {
	Kind       string   `json:"kind"`
	APIVersion string   `json:"apiVersion"`
	Items      []Gadget `json:"items"`
}
