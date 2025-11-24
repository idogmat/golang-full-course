package main

import "fmt"

func main() {
    x := 5
    fmt.Printf("x = %d в двоичном виде: %08b\n", x, x) // %08b дополнит нулями до 8 бит

    // Сдвиг влево
    result1 := x << 1
    fmt.Printf("5 << 1 = %d в двоичном виде: %08b\n", result1, result1)

    result11 := x << 2
    fmt.Printf("5 << 2 = %d в двоичном виде: %08b\n", result11, result11)

    result2 := x << 3
    fmt.Printf("5 << 3 = %d в двоичном виде: %08b\n", result2, result2)

    y := 80
    fmt.Printf("\ny = %d в двоичном виде: %08b\n", y, y)

    // Сдвиг вправо
    result3 := y >> 1
    fmt.Printf("80 >> 1 = %d в двоичном виде: %08b\n", result3, result3)

    result33 := y >> 2
    fmt.Printf("80 >> 2 = %d в двоичном виде: %08b\n", result33, result33)

    result4 := y >> 4
    fmt.Printf("80 >> 4 = %d в двоичном виде: %08b\n", result4, result4)
}