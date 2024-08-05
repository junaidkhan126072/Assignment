// orderedmap/orderedmap_test.go
package orderedmap

import (
	"testing"
)

// TestAdd tests the Add method of OrderedMap
func TestAdd(t *testing.T) {
	om := NewOrderedMap()

	// Test adding a new key-value pair
	om.Add("key1", "value1")
	value, exists := om.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Expected value1, got %s", value)
	}

	// Test updating an existing key-value pair
	om.Add("key1", "value2")
	value, exists = om.Get("key1")
	if !exists || value != "value2" {
		t.Errorf("Expected value2, got %s", value)
	}
}

// TestDelete tests the Delete method of OrderedMap
func TestDelete(t *testing.T) {
	om := NewOrderedMap()

	// Add a key-value pair
	om.Add("key1", "value1")

	// Test deleting an existing key
	om.Delete("key1")
	_, exists := om.Get("key1")
	if exists {
		t.Errorf("Expected key1 to be deleted")
	}

	// Test deleting a non-existing key (should not panic)
	om.Delete("key2")
}

// TestGet tests the Get method of OrderedMap
func TestGet(t *testing.T) {
	om := NewOrderedMap()

	// Test getting a non-existing key
	_, exists := om.Get("nonexistent")
	if exists {
		t.Errorf("Expected nonexistent key to not exist")
	}

	// Test getting an existing key
	om.Add("key1", "value1")
	value, exists := om.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Expected value1, got %s", value)
	}
}

// TestGetAll tests the GetAll method of OrderedMap
func TestGetAll(t *testing.T) {
	om := NewOrderedMap()

	// Test getting all key-value pairs from an empty map
	allItems := om.GetAll()
	if len(allItems) != 0 {
		t.Errorf("Expected 0 items, got %d", len(allItems))
	}

	// Test getting all key-value pairs after adding items
	om.Add("key1", "value1")
	om.Add("key2", "value2")

	allItems = om.GetAll()
	if len(allItems) != 2 {
		t.Errorf("Expected 2 items, got %d", len(allItems))
	}

	if allItems[0].Key != "key1" || allItems[0].Val != "value1" {
		t.Errorf("Expected key1:value1, got %s:%s", allItems[0].Key, allItems[0].Val)
	}

	if allItems[1].Key != "key2" || allItems[1].Val != "value2" {
		t.Errorf("Expected key2:value2, got %s:%s", allItems[1].Key, allItems[1].Val)
	}
}
