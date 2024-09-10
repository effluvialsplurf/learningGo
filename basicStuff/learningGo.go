package main

import (
	"fmt"
)

func exploringStructs() {
	// struct
	type Pokemon struct {
		name    string
		attacks []string
		hp      int
	}

	pichu := Pokemon{}
	pichu.name = "Pichu"
	pichu.hp = 31
	pichu.attacks = []string{"volt tackle", "tail whip"}

	// anonymous struct
	magicCard := struct {
		stats map[string]int
		name  string
		stars int
	}{
		stats: map[string]int{
			"attack":  300,
			"defense": 200,
		},
		name:  "Winged Kuriboh",
		stars: 1,
	}

	fmt.Printf("pokemon: %v, and then magicCard: %v\n", pichu, magicCard)
}

func main() {
	// fmt.Println("Finally, a git project AND a hello world file :)")
	// exploringStructs()
	//
	//fmt.Println("We are trying to get the bytes within a file.")
	//fmt.Println(fileLen("../projectMenu/menu.go"))
	//
	//helloPrefix := prefixer("Hello")
	//fmt.Println(helloPrefix("Bob"))
	//fmt.Println(helloPrefix("Maria"))
}
