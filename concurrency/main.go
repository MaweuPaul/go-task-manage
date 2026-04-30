package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var results []string

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			result := fmt.Sprintf("point %d calculated", id)

			mu.Lock() // lock before writing
			results = append(results, result)
			mu.Unlock() // unlock after writing
		}(i)
	}

	wg.Wait()
	fmt.Println("Results:", results)
}
