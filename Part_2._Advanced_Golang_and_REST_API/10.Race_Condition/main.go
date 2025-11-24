package main

import (
	"fmt"
	"sync"
)

// 1) Горутины (goroutines)

// Это лёгкие «зелёные» потоки, которые создаются и управляются рантаймом Go.
// Их может быть тысячи и миллионы.

// 2) Потоки ОС (threads)
// Рантайм Go использует M:N планировщик:
// M — количество потоков ОС (системных потоков), которые реально работают на ядрах.
// N — количество горутин.
// Горутины мультиплексируются поверх потоков ОС.

// 🧠 Зависимость от ядер
// ✔ Да, косвенная зависимость есть:
// На компьютере с 1 ядром одновременно по-настоящему выполняется только 1 поток ОС, а значит и 1 горутина (остальные переключаются).
// На компьютере с 8 ядрами одновременно могут выполняться до 8 потоков ОС, а значит до 8 горутин одновременно в реальном параллелизме.
// То есть параллелизм ограничен количеством ядер.

var number int = 0 // for see Race Condition

// var number atomic.Int64 // for solve Race Condition problem

var mtx sync.Mutex // way with mutex

func incrise(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// number.Add(1)
		// mtx.Lock()
		number++
		// mtx.Unlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(30)
	// Launch 30 goroutines to increment the number concurrently for reach limits
	for i := 1; i <= 30; i++ {
		go incrise(wg)
		fmt.Println("Goroutine launched:", i)
	}

	wg.Wait()

	// fmt.Printf("our number = %d\n", number.Load()) // 30 000
	fmt.Printf("our number = %d\n", number) // 30 000
}
