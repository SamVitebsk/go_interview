package main

import (
	"fmt"
	"log"
	"os"
)

// Counter представляет счетчик с целочисленным значением.
type Counter struct {
	Value int
}

// Print выводит значение счетчика в консоль.
func (c Counter) Print() {
	fmt.Println(c.Value)
}

// demoEvaluationAtRegistration демонстрирует, что аргументы и получатель (receiver) отложенного
// вызова вычисляются в момент регистрации defer, а не в момент фактического выполнения.
// Объяснение: В строке `defer c.Print()` фиксируется текущее состояние `c` ({Value: 123}).
// Ожидаемый вывод: 123. Последующая мутация `c.Value = 456` не влияет на уже скопированное для вызова defer значение.
func demoEvaluationAtRegistration() {
	c := Counter{Value: 123}
	defer c.Print() // Значение c ({Value: 123}) вычислено и зафиксировано прямо сейчас
	c.Value = 456
}

// demoClosure демонстрирует отложенный вызов через замыкание (closure).
// Объяснение: Замыкание захватывает переменную `c` по ссылке (адресу), поэтому в момент фактического
// выполнения отложенного вызова перед выходом из функции читается актуальное значение (456).
// Ожидаемый вывод: 456.
func demoClosure() {
	c := Counter{Value: 123}
	defer func() {
		c.Print() // Читает актуальное значение c в момент выполнения defer
	}()
	c.Value = 456
}

// demoLIFOOrder демонстрирует порядок выполнения нескольких defer в консоль.
// Объяснение: Вызовы defer помещаются в стек (LIFO — Last In, First Out).
// Поэтому последний зарегистрированный defer выполняется первым.
// Ожидаемый вывод: 3, 2, 1.
func demoLIFOOrder() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
}

// demoOsExit демонстрирует, что при вызове os.Exit defer не выполняется.
// Объяснение: os.Exit(code) немедленно завершает процесс через системный вызов ОС
// без раскрутки стека (stack unwinding), поэтому отложенные вызовы игнорируются.
func demoOsExit() {
	defer fmt.Println("ЭТОТ ТЕКСТ НЕ ВЫВЕДЕТСЯ: defer проигнорирован из-за os.Exit")
	fmt.Println("Вызов os.Exit(0): процесс завершается немедленно...")
	os.Exit(0)
}

// demoLogFatal демонстрирует, что при вызове log.Fatal defer не выполняется.
// Объяснение: функции log.Fatal* печатают сообщение и внутри вызывают os.Exit(1),
// поэтому стек не раскручивается и defer не срабатывает.
func demoLogFatal() {
	defer fmt.Println("ЭТОТ ТЕКСТ НЕ ВЫВЕДЕТСЯ: defer проигнорирован из-за log.Fatal")
	log.Fatal("Вызов log.Fatal: логирование ошибки и вызов os.Exit(1)...")
}

func main() {
	fmt.Println("=== 1. Вычисление значения в момент регистрации defer ===")
	demoEvaluationAtRegistration()

	fmt.Println("\n=== 2. Вычисление значения в момент выполнения defer (через замыкание) ===")
	demoClosure()

	fmt.Println("\n=== 3. Порядок выполнения нескольких defer (LIFO) ===")
	demoLIFOOrder()

	fmt.Println("\n=== 4. Завершение через os.Exit (defer не сработает) ===")
	demoOsExit()

	// Вызов demoLogFatal не будет достигнут из-за завершения процесса в demoOsExit.
	fmt.Println("\n=== 5. Завершение через log.Fatal (defer не сработает) ===")
	demoLogFatal()
}
