package main

import (
	"fmt"
)

func exploringStructs() {
	// struct
	type Pokemon struct {
		name    string
		hp      int
		attacks []string
	}

	pichu := Pokemon{}
	pichu.name = "Pichu"
	pichu.hp = 31
	pichu.attacks = []string{"volt tackle", "tail whip"}

	// anonymous struct
	magicCard := struct {
		name  string
		stars int
		stats map[string]int
	}{
		name:  "Winged Kuriboh",
		stars: 1,
		stats: map[string]int{
			"attack":  300,
			"defense": 200,
		},
	}

	fmt.Printf("pokemon: %v, and then magicCard: %v\n", pichu, magicCard)
}

func main() {
	fmt.Println("Finally, a git project AND a hello world file :)")
	exploringStructs()
}
