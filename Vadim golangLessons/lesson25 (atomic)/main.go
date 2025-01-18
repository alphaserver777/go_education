package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Главная функция программы
func main() {
	// AddMutex()  // Демонстрация использования мьютекса для безопасного увеличения счётчика
	// AddAtomic() // Демонстрация использования атомарных операций для увеличения счётчика

	// Примеры других атомарных операций
	// StoreLoadSwap()
	// compareAndSwap()
	atomicVal()
}

// AddMutex демонстрирует использование мьютекса (sync.Mutex) для защиты совместного ресурса
func AddMutex() {
	start := time.Now() // Засекаем время выполнения

	var (
		counter int64          // Счётчик, который будем увеличивать
		wg      sync.WaitGroup // Группа ожидания для синхронизации горутин
		mu      sync.Mutex     // Мьютекс для блокировки доступа к счётчику
	)

	wg.Add(1000) // Устанавливаем количество горутин, которые нужно дождаться

	// Запускаем 1000 горутин
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счётчик ожидания при завершении горутины

			mu.Lock()   // Блокируем доступ к разделяемому ресурсу
			counter++   // Увеличиваем счётчик
			mu.Unlock() // Разблокируем доступ
		}()
	}

	wg.Wait()                                                   // Ожидаем завершения всех горутин
	fmt.Println(counter)                                        // Выводим итоговое значение счётчика
	fmt.Println("With mutex:", time.Now().Sub(start).Seconds()) // Выводим время выполнения
}

// AddAtomic демонстрирует использование атомарных операций для увеличения счётчика
func AddAtomic() {
	start := time.Now() // Засекаем время выполнения

	var (
		counter int64          // Счётчик, который будем увеличивать
		wg      sync.WaitGroup // Группа ожидания для синхронизации горутин
	)

	wg.Add(1000) // Устанавливаем количество горутин

	// Запускаем 1000 горутин
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()              // Уменьшаем счётчик ожидания при завершении горутины
			atomic.AddInt64(&counter, 1) // Атомарное увеличение счётчика
		}()
	}

	wg.Wait()                                                    // Ожидаем завершения всех горутин
	fmt.Println(counter)                                         // Выводим итоговое значение счётчика
	fmt.Println("With atomic:", time.Now().Sub(start).Seconds()) // Выводим время выполнения
}

// StoreLoadSwap демонстрирует базовые атомарные операции: загрузка, запись и обмен значений
func StoreLoadSwap() {
	var counter int64 // Счётчик для демонстрации

	// atomic.LoadInt64 считывает текущее значение атомарно
	fmt.Println(atomic.LoadInt64(&counter)) // Выводим начальное значение (0)

	// atomic.StoreInt64 записывает новое значение атомарно
	atomic.StoreInt64(&counter, 5)          // Устанавливаем значение 5
	fmt.Println(atomic.LoadInt64(&counter)) // Выводим значение (5)

	// atomic.SwapInt64 атомарно заменяет значение и возвращает старое
	fmt.Println(atomic.SwapInt64(&counter, 10)) // Заменяем на 10 и выводим старое значение (5)
	fmt.Println(atomic.LoadInt64(&counter))     // Выводим текущее значение (10)
}

// compareAndSwap демонстрирует использование atomic.CompareAndSwapInt64
func compareAndSwap() {
	var (
		counter int64          // Счётчик
		wg      sync.WaitGroup // Группа ожидания
	)

	wg.Add(100) // Устанавливаем количество горутин

	// Запускаем 100 горутин
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done() // Уменьшаем счётчик ожидания при завершении горутины

			// atomic.CompareAndSwapInt64 проверяет, совпадает ли текущее значение с ожидаемым,
			// и если да, заменяет его новым
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return // Если значение уже изменено другой горутиной, выходим
			}

			fmt.Println("Swapped goroutine number is", i) // Если успешно заменили, выводим номер горутины
		}(i)
	}

	wg.Wait()            // Ожидаем завершения всех горутин
	fmt.Println(counter) // Выводим итоговое значение счётчика
}

// atomicVal демонстрирует использование atomic.Value для хранения любых типов данных
func atomicVal() {
	var (
		value atomic.Value // atomic.Value для хранения значений любого типа
	)

	value.Store(1)            // Сохраняем значение
	fmt.Println(value.Load()) // Загружаем и выводим сохранённое значение (1)

	fmt.Println(value.Swap(2)) // Заменяем значение на 2 и выводим старое (1)
	fmt.Println(value.Load())  // Выводим текущее значение (2)

	// Сравниваем текущее значение с ожидаемым и заменяем на новое, если совпадает
	fmt.Println(value.CompareAndSwap(2, 3)) // Заменяем 2 на 3, если текущее значение равно 2 (true)
	fmt.Println(value.Load())               // Выводим текущее значение (3)
}
