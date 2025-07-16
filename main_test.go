package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1", cargo: 42}
			et := &ElectricTruck{id: "2"}

			err := processTruck(nt)
			if err != nil {
				// Kill the program.
				t.Fatalf("Error processing truck: %s", err)
			}

			err = processTruck(et)
			if err != nil {
				// Kill the program.
				t.Fatalf("Error processing truck: %s", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("Normal truck cargo should be 0")
			}

			if et.battery != -2 {
				t.Fatal("Electric truck battery should be -2")
			}
		})
	})
}
