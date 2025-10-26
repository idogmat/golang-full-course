package main

import (
	"fmt"
)

func main() {

	score := 0
	for {
		fmt.Println("---------------------")

		score++

		if score == 15000 {
			break
		}

		fmt.Println(score, "score")

		// time.Sleep(500 * time.Millisecond)

	}
}
