package rboot

import (
	"fmt"
	"sync"
)

type Brain interface {
	Set(key string, value []byte) error
	Get(key string) []byte
	Remove(key string) error
}

var brains = make(map[string]func() Brain)

// 注册存储器
func RegisterBrain(name string, m func() Brain) {

	if name == "" {
		panic("RegisterBrain: brain must have a name")
	}
	if _, ok := brains[name]; ok {
		panic("RegisterBrain: brains named " + name + " already registered. ")
	}
	brains[name] = m
}

func DetectBrain(name string) (func() Brain, error) {
	if brain, ok := brains[name]; ok {
		return brain, nil
	}

	if len(brains) == 0 {
		return nil, fmt.Errorf("no Brain available")
	}

	if name == "" {
		if len(brains) == 1 {
			for _, brain := range brains {
				return brain, nil
			}
		}
		return nil, fmt.Errorf("multiple brains available; must choose one")
	}
	return nil, fmt.Errorf("unknown brain '%s'", name)
}

// memory brain
type memory struct {
	mu    sync.Mutex
	items map[string][]byte
}

// New constructs memory
func newMemory() Brain {
	return &memory{
		mu:    sync.Mutex{},
		items: make(map[string][]byte),
	}
}

// save ...
func (m *memory) Set(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.items[key] = value

	return nil
}

// find ...
func (m *memory) Get(key string) []byte {
	m.mu.Lock()
	defer m.mu.Unlock()

	v, ok := m.items[key]
	if !ok {
		return []byte{}
	}
	return v
}

// delete ...
func (m *memory) Remove(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.items, key)

	return nil
}

// register brain ...
func init() {
	RegisterBrain("memory", newMemory)
}
