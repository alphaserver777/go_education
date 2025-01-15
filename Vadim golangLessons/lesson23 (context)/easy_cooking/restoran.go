package main

import (
	"context"
	"fmt"
	"time"
)

// prepareDish имитирует процесс приготовления блюда
func prepareDish(ctx context.Context, dish string) {
	fmt.Println("Начинаем готовить:", dish)

	// Имитация долгой подготовки
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Приготовление блюда", dish, "отменено.")
			return // Останавливаем работу, если контекст отменён
		default:
			fmt.Printf("  Шаг %d приготовления %s...\n", i+1, dish)
			time.Sleep(1 * time.Second) // Имитация времени
		}
	}
	fmt.Println("Блюдо", dish, "готово!")

}

func main() {
	// 1. Создаем базовый контекст
	ctx := context.Background()

	// 2. Устанавливаем таймаут 3 секунды
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	//3. Запускаем горутину с приготовлением
	go prepareDish(ctxWithTimeout, "Стейк")
	time.Sleep(5 * time.Second) //Даем время горутине на выполнение
	fmt.Println("Программа завершена")

}
