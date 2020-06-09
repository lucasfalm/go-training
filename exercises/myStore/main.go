package main

func main() {
	store := list{
		book{name: "Book 1", value: 55},
		book{name: "Book 2", value: "250"},
		game{name: "Game 1", value: "100"},
	}

	for _, p := range store {
		p.print()
	}
}
