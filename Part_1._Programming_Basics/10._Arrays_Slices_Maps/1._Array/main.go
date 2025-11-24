package main

import "fmt"

func main() {
	intArr := [6]int{2, 4, 8, 16, 32, 64}
	animalArr := [6]string{"Dog", "Cat", "Bird", "Fish", "Horse", "Cow"}
	animalArr[0] = "Wolf"
	fmt.Println("arr:", intArr)
	fmt.Println("animalArr:", animalArr)
	// animalArr = append(animalArr, "Lion")
	fmt.Println("animalArr:", animalArr)

	// вывод на экран первым видом for
	for i := 0; i < len(intArr); i++ {
		fmt.Println(i, "->", intArr[i])
	}
	fmt.Println("---------------------")

	// умножаем каждый элемент на 2 первый раз первым видом for
	for i := 0; i < len(intArr); i++ {
		intArr[i] *= 2
	}

	// прибавляет каждому элементу, который больше 60, единицу вторым видом for
	for index, value := range intArr {
		if value > 60 {
			intArr[index] += 1

			// value++ не сработало бы
			// так как value это всего-лишь копия значения из массива
		}
	}

	// вывод на экран вторым видом for
	for index, value := range intArr {
		fmt.Println(index, "->", value)
	}
	fmt.Println("---------------------")

	// просто вывод всех индексов
	for index, _ := range intArr {
		fmt.Println("index:", index)
	}
	fmt.Println("---------------------")

	// просто вывод всех значений
	for _, value := range intArr {
		fmt.Println("value:", value)
	}
	fmt.Println("---------------------")
	spliceExamplr := make([]string, 5, 6)
	fmt.Printf("animalArr:ptr = %p length = %d capacity = %d\n", &animalArr, len(animalArr), cap(animalArr))
	fmt.Printf("spliceExamplr:ptr = %p length = %d capacity = %d\n", &spliceExamplr, len(spliceExamplr), cap(spliceExamplr))
	fmt.Println("---------------------")
}
