package adapter

import (
	"fmt"
	"sort"
)

type Registry struct {
	byName map[string]Adapter
}

func NewRegistry(adapters ...Adapter) (*Registry, error) {
	r := &Registry{byName: make(map[string]Adapter, len(adapters))}
	for _, a := range adapters {
		if a == nil {
			return nil, fmt.Errorf("adapter: nil adapter")
		}
		name := a.Name()
		if name == "" {
			return nil, fmt.Errorf("adapter: empty adapter name")
		}
		if _, exists := r.byName[name]; exists {
			return nil, fmt.Errorf("adapter: duplicate adapter %q", name)
		}
		r.byName[name] = a
	}
	return r, nil
}

func (r *Registry) Get(name string) (Adapter, bool) {
	a, ok := r.byName[name]
	return a, ok
}

func (r *Registry) All() []Adapter {
	names := make([]string, 0, len(r.byName))
	for name := range r.byName {
		names = append(names, name)
	}
	sort.Strings(names)

	out := make([]Adapter, 0, len(names))
	for _, name := range names {
		out = append(out, r.byName[name])
	}
	return out
}
