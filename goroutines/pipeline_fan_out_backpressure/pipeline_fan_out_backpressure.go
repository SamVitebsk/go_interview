package pipeline_fan_out_backpressure

import (
	"context"
	"sync"
)

// User представляет модель данных пользователя.
type User struct {
	ID     int
	Active bool
}

// gen последовательно отправляет идентификаторы из ids в выходной канал с проверкой отмены контекста.
// Роль в конвейере: Generator (источник данных).
// Backpressure: небуферизированный канал блокирует отправку до готовности downstream-воркеров,
// предотвращая генерацию данных быстрее, чем конвейер успевает их обрабатывать.
// Канал закрывается отправителем при завершении работы.
func gen(ctx context.Context, ids []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, id := range ids {
			select {
			case <-ctx.Done():
				return
			case out <- id:
			}
		}
	}()

	return out
}

// load параллельно загружает данные пользователей по ID.
// Паттерны:
// - Fan-Out: 5 горутин-воркеров параллельно разбирают задачи из одного входного канала ids.
// - Fan-In: все 5 воркеров мультиплексируют свои результаты в один общий выходной канал users.
// - Backpressure: фиксированный пул воркеров (5) и буфер канала (5) жестко ограничивают объем данных в полете.
// Закрытие общего канала координируется отдельной горутиной через sync.WaitGroup после завершения всех воркеров.
func load(ctx context.Context, ids <-chan int) <-chan User {
	workersCount := 5
	users := make(chan User, workersCount)

	var wg sync.WaitGroup
	wg.Add(workersCount)

	// Fan-Out: запуск пула воркеров
	for range workersCount {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case id, ok := <-ids:
					if !ok {
						return
					}

					user := User{ID: id, Active: id%2 == 0}
					select {
					case <-ctx.Done():
						return
					case users <- user: // Fan-In: запись в общий канал
					}
				}
			}
		}()
	}

	// Координация закрытия канала после завершения Fan-In
	go func() {
		wg.Wait()
		close(users)
	}()

	return users
}

// filter отбирает из входного канала только активных пользователей (Active == true).
// Роль в конвейере: Transformer / Filter (стадия 1-к-1).
// Backpressure: небуферизированный канал res передает противодавление обратно: если collect
// не успевает читать, filter блокируется на записи и перестает вычитывать users, блокируя load.
// Проверяет отмену контекста и на вычитке, и на отправке (внутренний select) для защиты от утечек горутин.
func filter(ctx context.Context, users <-chan User) <-chan User {
	res := make(chan User)

	go func() {
		defer close(res)
		for {
			select {
			case <-ctx.Done():
				return
			case u, ok := <-users:
				if !ok {
					return
				}

				if u.Active {
					select {
					case <-ctx.Done():
						return
					case res <- u:
					}
				}
			}
		}
	}()

	return res
}

// collect синхронно вычитывает отфильтрованных пользователей из канала и собирает их в итоговый срез []User.
// Роль в конвейере: Sink / Consumer (терминальная стадия).
// Не порождает фоновых горутин, блокирует вызывающий поток до завершения конвейера или отмены контекста.
func collect(ctx context.Context, users <-chan User) []User {
	var res []User

	for {
		select {
		case <-ctx.Done():
			return res
		case user, ok := <-users:
			if !ok {
				return res
			}

			res = append(res, user)
		}
	}
}
