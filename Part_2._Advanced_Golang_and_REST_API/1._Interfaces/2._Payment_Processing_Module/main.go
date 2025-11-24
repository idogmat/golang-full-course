package main

import (
	"interfaces/car"
	"interfaces/payments"
	"interfaces/payments/methods"
	"time"

	"github.com/k0kubun/pp"
)

func main() {
	before := time.Now()
	// Выбираем метод оплаты
	// method := methods.NewBank()   // Оплата через банк
	// method := methods.NewPayPal() // Оплата через PayPal
	method := methods.NewCrypto() // Оплата через крипто-кошелёк

	// Создаём модуль проведения оплат
	paymentModule := payments.NewPaymentModule(method)

	// Проводим 3 оплаты
	paymentModule.Pay("Бургер", 5)
	phonePaymentID := paymentModule.Pay("Телефон", 500)
	paymentModule.Pay("Игра", 20)

	// По сохранённому ID оплаты за телефон, мы отменяем эту самую оплату
	paymentModule.Cancel(phonePaymentID)

	// Получаем информацию по всем проведённым оплатам
	allInfo := paymentModule.AllInfo()

	// Выводим эту информацию в консоль
	pp.Println("Информация по всем проведённым оплатам:", allInfo)

	Toyota := car.NewCarModule(car.CarCharacteristics{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2020,
		Color: "Синий",
	})

	pp.Println("Информация по автомобилю:", Toyota)
	Toyota.ChangeColor("Red")
	Toyota.Drive("Поездка на работу", 15)
	pp.Println("Информация по автомобилю:", Toyota)

	after := time.Now()
	pp.Println("Время выполнения Sub:", after.Sub(before))
	pp.Println("Время выполнения Since:", time.Since(before))
	pp.Println("Время выполнения Since:", time.Since(before).Microseconds())
}
