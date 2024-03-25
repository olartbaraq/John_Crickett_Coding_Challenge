package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("incomplete arguments")
	}

	if os.Args[1] != "sort" {
		log.Fatalf("%v is incorrect for the first argument", os.Args[1])
	}

	if len(os.Args) == 3 {
		file, err := os.Open(os.Args[2])

		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		//Decleare a slice to keep all the words
		var Arraywords []string

		for scanner.Scan() {
			Arraywords = append(Arraywords, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		n := len(Arraywords)

		// Apply Bubble sort algorthm to sort each words in the array

		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if Arraywords[i] > Arraywords[j] {
					temp := Arraywords[i]
					Arraywords[i] = Arraywords[j]
					Arraywords[j] = temp
				}
			}
		}

		newFile, err := os.Create("sorted_words.txt")

		for _, words := range Arraywords {
			_, err := newFile.WriteString(words + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}

		defer newFile.Close()

	}
}
