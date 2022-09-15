package dmap

type DeterministicMap[K comparable, V any] struct {
	data map[K]V
	keys []K
}

func NewMap[K comparable, V any]() *DeterministicMap[K, V] {
	data := make(map[K]V)
	keys := make([]K, 0)

	return &DeterministicMap[K, V]{
		data: data,
		keys: keys,
	}
}

func (m *DeterministicMap[K, V]) Get(key K) (V, bool) {
	v, found := m.data[key]
	return v, found
}

func (m *DeterministicMap[K, V]) Set(key K, val V) {
	if _, found := m.data[key]; !found {
		m.data[key] = val
		m.keys = append(m.keys, key)
	}
}

func (m *DeterministicMap[K, V]) Range() []K {
	return m.keys
}
