package main

import (
	"fmt"
	"sync"
)

// Inventory Struct
type Inventory struct {
	items map[string]int
	mu    sync.Mutex
}

// NewInventory Initiate new inventory
func NewInventory() *Inventory {
	return &Inventory{
		items: make(map[string]int),
	}
}

// AddToInventory add inventory to struct
func (inv *Inventory) AddToInventory(item string, quantity int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	inv.items[item] += quantity
}

// RemoveFromInventory delete data from inventory
func (inv *Inventory) RemoveFromInventory(item string, quantity int) {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if currentQuantity, ok := inv.items[item]; ok {
		if currentQuantity >= quantity {
			inv.items[item] -= quantity
		} else {
			fmt.Println(fmt.Sprintf("Error: %s Insufficient quantity in inventory.", item))
		}
	} else {
		fmt.Println(fmt.Sprintf("Error: %s Item not found in inventory.", item))
	}
}

// GetInventory Get Inventory
func (inv *Inventory) GetInventory() map[string]int {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	copy := make(map[string]int)
	for item, quantity := range inv.items {
		copy[item] = quantity
	}
	return copy
}

func main() {
	inv := NewInventory()

	// Simulate inventory transactions
	inv.AddToInventory("Desktop", 5)
	inv.RemoveFromInventory("Laptop", 4)
	inv.RemoveFromInventory("Desktop", 2)

	// Print current inventory
	fmt.Println("Current Inventory:", inv.GetInventory())
}
