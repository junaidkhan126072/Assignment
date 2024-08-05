// server/server_test.go
package server

import (
	"testing"

	"os"
	"strings"

	"github.com/assignment/utils"
)

func TestHandleAddItem(t *testing.T) {
	s := NewServer()

	s.handleAddItem("key1", "value1")
	value, exists := s.orderedMap.Get("key1")
	if !exists || value != "value1" {
		t.Errorf("Expected value1, got %s", value)
	}
}

func TestHandleDeleteItem(t *testing.T) {
	s := NewServer()

	s.handleAddItem("key1", "value1")
	s.handleDeleteItem("key1")
	_, exists := s.orderedMap.Get("key1")
	if exists {
		t.Errorf("Expected key1 to be deleted")
	}
}

func TestHandleGetItem(t *testing.T) {
	// Prepare the output file
	utils.OutputFile = "test_output.txt"

	s := NewServer()

	// Test getting an existing item
	s.handleAddItem("key1", "value1")
	s.handleGetItem("key1")

	content, err := os.ReadFile(utils.OutputFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !strings.Contains(string(content), "Key: key1, Value: value1") {
		t.Errorf("Expected output to contain 'Key: key1, Value: value1', got %s", string(content))
	}

	// Test getting a non-existing item
	s.handleGetItem("key2")
	content, err = os.ReadFile(utils.OutputFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !strings.Contains(string(content), "Key: key2 not found") {
		t.Errorf("Expected output to contain 'Key: key2 not found', got %s", string(content))
	}
	defer os.Remove(utils.OutputFile)

}

func TestHandleGetAllItems(t *testing.T) {
	// Prepare the output file
	utils.OutputFile = "test_output_all.txt"

	s := NewServer()

	s.handleAddItem("key1", "value1")
	s.handleAddItem("key2", "value2")

	s.handleGetAllItems()
	content, err := os.ReadFile(utils.OutputFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !strings.Contains(string(content), "Key: key1, Value: value1") {
		t.Errorf("Expected output to contain 'Key: key1, Value: value1', got %s", string(content))
	}

	if !strings.Contains(string(content), "Key: key2, Value: value2") {
		t.Errorf("Expected output to contain 'Key: key2, Value: value2', got %s", string(content))
	}
	defer os.Remove(utils.OutputFile)

}
