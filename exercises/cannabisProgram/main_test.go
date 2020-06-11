package main

import (
	"testing"
)

func setup(t *testing.T) *cannabis {
	cExample := cannabis{
		name: "Test",
	}
	return &cExample
}
func TestUpdateFlower(t *testing.T) {
	type test struct {
		f    *cannabis
		fail bool
	}
	f := setup(t)

	tests := []test{
		test{f, false},
	}

	for _, t := range tests {
		_ = t
	}
	nF := flowers{
		name: "Test",
		thc:  35,
		geo:  make(map[string]int), // Initializing map
	}
	eX := map[string]int{
		"test": 10,
	}
	err := f.updateFlower("Test", 35, eX)

	if err != nil {
		t.Error(f)
	}

	if f.flowers[0].name != nF.name {
		t.Error(f)
	}
}
