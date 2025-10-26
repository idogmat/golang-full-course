package main

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

func main() {
		var i int = 10
		var i2 int64 = 10
    var s string = "Hello Hello Hello Hello Hello Hello"
    var sl []int = []int{1, 2, 3}

    // Размер примитивных типов и дескрипторов
    fmt.Println("Sizeof(int):", unsafe.Sizeof(i))   // 8 (на 64-битной системе)
    fmt.Println("Sizeof(int64):", unsafe.Sizeof(i2))  // 1
    fmt.Println("Sizeof(string):", unsafe.Sizeof(s)) // 16 (указатель + длина)
    fmt.Println("Sizeof([]int):", unsafe.Sizeof(sl)) 
	// это пример использования переменных
	var j int = 1
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	number := 15
	text := "Hello"
	drob := 11.25
	boolean := true

	fmt.Println("number:", number)
	fmt.Println("text:", text)
	fmt.Println("drob:", drob)
	fmt.Println("boolean:", boolean)

	var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
