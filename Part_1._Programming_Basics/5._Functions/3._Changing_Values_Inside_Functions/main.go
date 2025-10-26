package main

import "fmt"

func main() {
	number := 5
	text := "hello"

	fmt.Println("number до:", number)
	fmt.Println("text до:", text)

	foo(number, text)

	fmt.Println("number после:", number, &number)
	fmt.Println("text после:", text, &text)
	
	foo2(&number, &text)

	fmt.Println("number после foo2 передающая указатель:", number)
	fmt.Println("text после foo2 передающая указатель:", text)

}

func foo(n int, t string) {
	n = 10
	t = "world"
}

func foo2(n *int, t *string) { // Объявление функции foo2, которая принимает два аргумента: указатель на int и указатель на string
	fmt.Println("n адрес:", n, "значение:", *n)
	fmt.Println("t адрес:", t, "значение:", *t)
	*n = 10                    // Разыменование указателя n и присваивание значения 10 переменной, на которую он указывает
	*t = "world"               // Разыменование указателя t и присваивание значения "world" переменной, на которую он указывает
}
