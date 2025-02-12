package main

import (
	"errors"
	"fmt"
)

type Item struct {
	name     string
	quantity int
}

type Inventory map[string]Item

const MaxInventorySize = 10

func addItem(inv Inventory, name string, quantity int) error {

	if _, exists := inv[name]; exists {
		return fmt.Errorf("Item %s already exists", name)
	}

	if len(inv) >= MaxInventorySize {
		return &InventoryFullError{
			inventorySize: MaxInventorySize,
		}
	}

	inv[name] = Item{name: name, quantity: quantity}

	return nil
}

func wrapError(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}

func printErrorChain(err error) {
	if err == nil {
		return
	}

	fmt.Printf("Error: %v\n", err)
	printErrorChain(errors.Unwrap(err))
}

func removeItem(inv Inventory, name string) error {
	if _, exists := inv[name]; !exists {
		return wrapError(errors.New("Item not exists"), name)
	}

	return nil
}

type InventoryFullError struct {
	inventorySize int
}

func (e InventoryFullError) Error() string {
	return fmt.Sprintf("Max inventory size %d", e.inventorySize)
}

func updateItemQuantity(inv Inventory, name string, quantity int) {
	if _, exists := inv[name]; !exists {
		panic(fmt.Sprintf("%s non exists", name))
	}

	item := inv[name]
	item.quantity = quantity
	inv[name] = item
}

func main() {
	inventory := make(Inventory)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if err := addItem(inventory, "apple", 5); err != nil {
		fmt.Printf("Error adding apple: %v\n", err)
	}

	if err := addItem(inventory, "banana", 7); err != nil {
		fmt.Printf("Error adding banana: %v\n", err)
	}

	fmt.Println("\nRemoving non-existent item from inventory: orange")
	if err := removeItem(inventory, "orange"); err != nil {
		printErrorChain(err)
	}

	updateItemQuantity(inventory, "apple", 10)

	for i := range 9 {
		if err := addItem(inventory, string(i), 1); err != nil {
			fmt.Printf("Error adding new item: %v\n", err)
		}
	}

	fmt.Println("\nAttempting to update non-existent item: orange")
	updateItemQuantity(inventory, "orange", 5)
}
