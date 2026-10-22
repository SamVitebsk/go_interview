package main

import "fmt"

/*
Задание:

1. Вывод и почему answer = 41:
Вывод: base -> init 41 -> main 41.
Причина: answer зависит от base, поэтому компилятор сначала вычисляет base (печать "base", возврат 40),
затем вычисляет answer (40 + 1 = 41), после чего запускает init(), и в конце — main().

2. Если поменять местами base и answer:
Результат не изменится: логический граф зависимостей останется прежним (answer по-прежнему зависит от base).

3. Пример циклической зависимости (ошибка компиляции):
// var a = b + 1
// var b = a + 1
// Ошибка: initialization cycle for a

4. Шаги запуска сервиса, выполняемые явно в main:
- Парсинг конфига и флагов (config.Load)
- Подключение к БД/брокерам с context.Context (db.Connect)
- Применение миграций
- DI / сборка слоев (репозитории, сервисы, хендлеры)
- Запуск HTTP/gRPC серверов
- Graceful shutdown по сигналам ОС (SIGINT/SIGTERM)
В init() их выносить нельзя из-за отсутствия error, context и невозможности мокать в тестах.
*/

var answer = base + 1
var base = initialBase()

func initialBase() int {
	fmt.Println("base")
	return 40
}

func init() {
	fmt.Println("init", answer)
}

func main() {
	fmt.Println("main", answer)
}
