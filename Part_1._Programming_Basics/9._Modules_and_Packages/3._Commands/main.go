package main

import (
	"fmt"
	"simple-example/utils"
)

func main() {
	utils.InfoLog("Message from Logger")
	fmt.Println("run for create go.mod: go mod init simple-example")
	fmt.Println("run for refine go.mod: go mod tidy")
	fmt.Println(utils.Add(2, 3))
	fmt.Println(utils.Subtract(5, 2))
	fmt.Println(utils.Multiply(3, 4))
	fmt.Println(utils.Divide(8, 2))
	fmt.Println(utils.Divide(8, 0))
	fmt.Println("ready")
}
