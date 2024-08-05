// orderedmap/orderedmap.go
package orderedmap

import (
	"container/list"
	"sync"
)

// KeyValue represents a key-value pair
type KeyValue struct {
	Key string
	Val string
}

// OrderedMap represents an ordered map data structure
type OrderedMap struct {
	data map[string]*list.Element
	list *list.List
	mu   sync.RWMutex
}

// NewOrderedMap creates and returns a new ordered map
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		data: make(map[string]*list.Element),
		list: list.New(),
	}
}

// Add adds a key-value pair to the ordered map
func (om *OrderedMap) Add(key, val string) {
	om.mu.Lock()
	defer om.mu.Unlock()
	if elem, exists := om.data[key]; exists {
		elem.Value.(*KeyValue).Val = val
	} else {
		kv := &KeyValue{Key: key, Val: val}
		om.data[key] = om.list.PushBack(kv)
	}
}

// Delete removes a key-value pair from the ordered map
func (om *OrderedMap) Delete(key string) {
	om.mu.Lock()
	defer om.mu.Unlock()
	if elem, exists := om.data[key]; exists {
		om.list.Remove(elem)
		delete(om.data, key)
	}
}

// Get retrieves a value by key
func (om *OrderedMap) Get(key string) (string, bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()
	if elem, exists := om.data[key]; exists {
		return elem.Value.(*KeyValue).Val, true
	}
	return "", false
}

// GetAll returns all key-value pairs in insertion order
func (om *OrderedMap) GetAll() []KeyValue {
	om.mu.RLock()
	defer om.mu.RUnlock()
	var result []KeyValue
	for elem := om.list.Front(); elem != nil; elem = elem.Next() {
		kv := elem.Value.(*KeyValue)
		result = append(result, *kv)
	}
	return result
}
