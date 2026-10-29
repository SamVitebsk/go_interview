package deadlock_goroutine_leak_close_channel

import (
	"context"
)

// Numbers генерирует последовательность чисел, управляемую через context.Context
// При отмене контекста горутина завершается, а канал безопасно закрывается.
//
// Почему старая версия протекала: небуферизованная запись out <- i блокировалась навсегда,
// когда получатель прекращал чтение, так как не было сигнала отмены (ctx.Done),
// из-за чего зависшая горутина пожизненно оставалась в памяти (goroutine leak).
func Numbers(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()

	return out
}
