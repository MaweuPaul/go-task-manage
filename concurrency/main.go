package main

import (
	"fmt"
)

// interfaces
type CoordinateSystem interface {
	Convert(x, y float64) (float64, float64)
	Name() string
}

type CassiniConverter struct {
	FalseEasting  float64
	FalseNorthing float64
}
type UTMConverter struct {
	Easting  float64
	Northing float64
	Zone     int
}

func (c CassiniConverter) Convert(x, y float64) (float64, float64) {
	return x + c.FalseEasting, y + c.FalseNorthing
}
func (c CassiniConverter) Name() string {
	return "Cassini-Soldner"
}

func (u UTMConverter) Convert(x, y float64) (float64, float64) {
	return x + float64(u.Zone), y + float64(u.Zone)
}

func (u UTMConverter) Name() string {
	return fmt.Sprintf("UTM Zone %d", u.Zone)
}

func convertAndPrint(cs CoordinateSystem, x, y float64) {
	name := cs.Name()
	newX, newY := cs.Convert(x, y)
	fmt.Printf("%s: (%.2f, %.2f) → (%.2f, %.2f)\n", name, x, y, newX, newY)
}

func main() {
	// channels
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

	utm := UTMConverter{Easting: 500000, Northing: 9000000, Zone: 37}
	cassini := CassiniConverter{FalseEasting: 100, FalseNorthing: 200}

	convertAndPrint(utm, 100, 200)
	convertAndPrint(cassini, 100, 200)

}
