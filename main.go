package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// Create your own error messages.
var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

// Blueprint to define methods.
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

// Implement interfaces.
func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Started processing truck %+v\n", truck)

	// Simulate some processing time.
	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		// Context for debugging. Alternative to 'new()'.
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		// Context for debugging. Alternative to 'new()'.
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finished processing truck %+v\n", truck)

	return nil
}

// processFleet demonstrates concurrent processing of multiple trucks.
func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)

		go func(t Truck) {
			if err := processTruck(t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait()

	return nil
}

func main() {
	// Processing all at the same time.
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100},
	}

	// Process all trucks concurrently
	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed successfully!")
}
