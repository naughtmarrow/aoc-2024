package main

import (
	"aoc/day2/stars"
	"fmt"
)

func main() {
	safe := stars.First()
	fmt.Println("The amount of safe codes in part1 is: ", safe)

	safe = stars.Second()
	fmt.Println("The amount of safe codes in part2 is: ", safe)
}
