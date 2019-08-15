// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// person represents a person.
type person struct {
	name   string
	weight int     // kg
	height float64 // meter
}

// bmi calculates the body mass index of player.
func (p *person) bmi() float64 {
	return float64(p.weight) / (p.height * p.height)
}

func main() {

	// Create a few players.
	ps := []person{
		{"ada", 79, 1.80},
		{"bob", 150, 2.0},
		{"dan", 55, 1.65},
	}

	// Display the batting average for each player.
	for i := range ps {
		fmt.Printf("%s: bmi=%2.2f\n", ps[i].name, ps[i].bmi())
	}

	// Why did I not choose this form?
	for _, p := range ps {
		fmt.Printf("%s: bmi=%2.2f\n", p.name, p.bmi())
	}
}
