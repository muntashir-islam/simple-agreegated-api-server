package main

type MemoryStore struct {
	items []Gadget
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{items: []Gadget{}}
}

func (s *MemoryStore) List() GadgetList {
	return GadgetList{
		Kind:       "GadgetList",
		APIVersion: "gadgets.muntashirislam.com/v1",
		Items:      s.items,
	}
}

func (s *MemoryStore) Create(g Gadget) {
	g.Kind = "Gadget"
	g.APIVersion = "gadgets.muntashirislam.com/v1"
	s.items = append(s.items, g)
}
