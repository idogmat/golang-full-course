// /C:/PROJECTS/golang_starter/golang-full-course/Part_1._Programming_Basics/9._Modules_and_Packages/4._Buffer/main.go
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Создаём буфер (bytes.Buffer реализует io.Reader и io.Writer)
	var buf bytes.Buffer
 	fmt.Println("stdout fd:", os.Stdout.Fd())
	// Дескриптор (file descriptor) — это целочисленный идентификатор
	// открытого ресурса в операционной системе (файла, сокета, пайпа и т.п.). Коротко:
 	// В Unix-подобных системах дескриптор — небольшое целое (0,1,2,...), которое ядро использует для доступа к открытым объектам.
	// Записываем данные в буфер
	buf.WriteString("Hello, buffer!\n")
	buf.Write([]byte("Second line\nThird line\n"))

	fmt.Println("buffer length before reading:", buf.Len())

	// Читаем по строкам с помощью bufio.Reader (чтение потребляет содержимое bytes.Buffer)
	r := bufio.NewReader(&buf)
	for {
		line, err := r.ReadString('\n')
		// Выводим строку (если строка без '\n' — она всё равно будет выведена при EOF)
		fmt.Print("line: ", line)

		if err == io.EOF {
			// достигнут конец — последний кусок уже напечатан, выходим
			break
		}
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
	}

	fmt.Println("buffer length after reading:", buf.Len()) // обычно 0, т.к. чтение съело буфер

	// Подготовим новый содержимый для демонстрации io.Copy
	buf.Reset()
	buf.WriteString("Final message via io.Copy\n")

	// Копируем содержимое буфера в stdout (и тем самым опустошаем буфер)
	if _, err := io.Copy(os.Stdout, &buf); err != nil {
		fmt.Println("copy error:", err)
	}
	// После копирования буфер пуст
	fmt.Println("\nbuffer length after copy:", buf.Len())
	
}