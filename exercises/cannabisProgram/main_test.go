package main

import (
	"fmt"
	"testing"
)

type cannabisFixture struct {
	testName     string
	expectError  bool
	errorMessage string

	mockCannabis *cannabis
}

func (c *cannabisFixture) setup(t *testing.T) {
	c.mockCannabis = newCannabisMock()
}

func TestUpdateFlower(t *testing.T) {
	table := []cannabisFixture{
		{
			testName:    "Create new flower",
			expectError: false,
		},
		{
			testName:    "Create new flower with error",
			expectError: true,
		},
	}

	for index, f := range table {
		t.Run(fmt.Sprintf("%v-%v", index, f.testName), func(t *testing.T) {
			f.setup(t)

			eX := map[string]int{
				"Brazil": 1500,
			}

			err := f.mockCannabis.updateFlower("Test", 35, eX)
			if !f.expectError {
				if err != nil {
					t.Fatal("an error was not expected, but occurred")
				}
			}

			err = f.mockCannabis.updateFlower("Test", -35, eX)
			if f.expectError {
				if err == nil {
					t.Fatal("an error was expected, but not occurred")
				}
			}
		},
		)
	}
}

func newCannabisMock() *cannabis {
	return &cannabis{
		name: "Sativa",
	}
}
