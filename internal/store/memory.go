package store

import (
	"mws-test/internal/api"
	"sync"
)

type MemoryStore struct {
	mu     sync.RWMutex
	nextID int64
	cats   map[int64]api.Cat
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		cats: make(map[int64]api.Cat),
	}
}

func (s *MemoryStore) List() []api.Cat {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]api.Cat, 0, len(s.cats))
	for _, c := range s.cats {
		out = append(out, c)
	}
	return out
}

func (s *MemoryStore) Create(c api.NewCat) api.Cat {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	newCat := api.Cat{
		ID:    s.nextID,
		Name:  c.Name,
		Age:   c.Age,
		Color: c.Color,
	}
	s.cats[s.nextID] = newCat
	return newCat
}

func (s *MemoryStore) Get(id int64) (api.Cat, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.cats[id]
	return c, ok
}

func (s *MemoryStore) Update(id int64, upd api.UpdateCat) (api.Cat, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.cats[id]
	if !ok {
		return api.Cat{}, false
	}
	updated := api.Cat{
		ID:    id,
		Name:  upd.Name,
		Age:   upd.Age,
		Color: upd.Color,
	}
	s.cats[id] = updated
	return updated, true
}

func (s *MemoryStore) Delete(id int64) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.cats[id]
	if ok {
		delete(s.cats, id)
	}
	return ok
}
