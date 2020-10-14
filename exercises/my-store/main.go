package main

func main() {
	store := list{
		&book{product{name: "Book 1", value: 55}},
		&book{product{name: "Book 2", value: "250"}},
		&game{product{name: "Game 1", value: "100"}},
	}

	for _, p := range store {
		p.print()
		p.discount(10) // Apply 10% discount
	}
}
