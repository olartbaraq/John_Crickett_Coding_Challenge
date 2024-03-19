package tests

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountRune(t *testing.T) {
	file, err := os.Open("../test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("test.txt does not exist")
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

	// implementing hash map to keep count of each runes
	for scanner.Scan() {

		char := scanner.Text()[0] // Get the rune from the scanner
		hashmap[string(char)]++
	}

	for _, count := range hashmap {
		assert.NotZero(t, count)
	}

}
