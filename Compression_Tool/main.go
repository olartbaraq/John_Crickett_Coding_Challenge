package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Incomplete arguments: the number of commands must be 2")
	}

	fileText := os.Args[1]

	file, err := os.Open(fileText)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("%v does not exist", fileText)
		}
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	hashmap := make(map[string]int)

	fmt.Print(hashmap)

	// implementing hash map to keep count of each runes
	for scanner.Scan() {

		char := scanner.Text()[0] // Get the rune from the scanner
		hashmap[string(char)]++
	}

	for char, count := range hashmap {
		fmt.Printf("%v: %v\n", char, count)
	}
}
