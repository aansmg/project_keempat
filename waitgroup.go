package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "mango", "banana", "rambutan"}
	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go PrintFruits(index, fruit, &wg)
	}
	wg.Wait()
}

func PrintFruits(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index :=> %d, fruit => %s\n", index, fruit)
	wg.Done()
}
