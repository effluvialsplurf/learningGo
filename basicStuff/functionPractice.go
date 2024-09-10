package main

import (
	"log"
	"os"
)

func fileLen(input string) (int, error) {
	// Open a file of name: input
	file, err := os.Open(input)
	// check for errors
	if err != nil {
		log.Fatal(err)
	}
	// close the file after we are done using it
	defer file.Close()

	// read the contents of the file
	contents, err := os.ReadFile(input)
	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	return len(contents), err
}

func prefixer(toPre string) func(string) string {
	return func(pre string) string {
		return toPre + " " + pre
	}
}
