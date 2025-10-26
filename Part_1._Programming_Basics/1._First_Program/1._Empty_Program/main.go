package main

import (
	"fmt"
	"math"
)


func main() {
	var example string
	example = "test"
	fmt.Printf("example: %s\n", example)

	score := 0
	score++
	fmt.Println(score)
	score++
	fmt.Println(score)
	fmt.Println(`score:`, score)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(float64(score)))
	// это пример использования переменных
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
