package storage

import "sync"

// in memory implementation of engine
type MemoryStore struct {
	mu sync.RWMutex // Standard Read write mutex
	data map[string][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string][]byte),
	}
}

func (m *MemoryStore) Put(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// copy the data to avoid shared memory
	valTemp := make([]byte, len(value))
	copy(valTemp, value)

	m.data[key] = valTemp
	return nil
}

func (m *MemoryStore) Get(key string) ([]byte, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	val, exists := m.data[key]
	if !exists {
		return nil, ErrKeyNotFound
	}

	valTemp := make([]byte, len(val))
	copy(valTemp, val)

	return valTemp, nil
}

func (m *MemoryStore) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, key)
	return nil
}