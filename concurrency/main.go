package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	var results []string

	for i := 1; i <= 5; i++ {
		go func(id int) {
			result := fmt.Sprintf("Result from goroutine %d", id)
			ch <- result
		}(i)
	}

	for i := 0; i < 5; i++ {
		results = append(results, <-ch)
	}
	fmt.Println("Results:", results)
}
