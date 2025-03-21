package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// NOTE:Wait group
	// withoutWait()
	// withWait()
	wrongAdd()

	// Mutex
	// writeWithoutConcurrent()
	// writeWithoutMutex()
	// writeWithMutex()
	// readWithMutex()
	//readWithRWMutex()
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup
	/*
		Мы создаем переменную wg типа sync.WaitGroup. WaitGroup используется для ожидания завершения группы горутин (goroutines). Это полезно, когда у вас есть несколько горутин, и вы хотите дождаться, пока все они завершат свою работу, прежде чем продолжить выполнение программы.
	*/
	wg.Add(10)
	/*
		Мы вызываем метод Add(10) на переменной wg. Это добавляет 10 ожидаемых горутин в группу. Это означает, что мы ожидаем, что 10 горутин завершат свою работу, прежде чем мы продолжим выполнение программы.
	*/

	for i := 0; i < 10; i++ {
		go func(i int) {
			//defer wg.Done()
			fmt.Println(i + 1)
			wg.Done()
		}(i)
		/*
		   Здесь мы запускаем цикл, который выполняется 10 раз. На каждой итерации цикла мы запускаем новую горутину с помощью ключевого слова go.

		   Горутина: Это легковесный поток, который выполняется параллельно с другими горутинами. Внутри горутины мы передаем текущее значение i в анонимную функцию.

		   Анонимная функция: Это функция, которая не имеет имени и выполняется сразу после создания. Она принимает параметр i и выводит его значение, увеличенное на 1, с помощью fmt.Println(i + 1).

		   wg.Done(): После того как горутина завершает свою работу (в данном случае, после вывода числа), она вызывает wg.Done(), чтобы уменьшить счетчик WaitGroup на 1.
		*/
	}

	wg.Wait()
	/*
	   wg.Wait(): Этот метод блокирует выполнение программы до тех пор, пока счетчик WaitGroup не станет равным 0. Это означает, что мы ждем, пока все 10 горутин завершат свою работу.
	*/
	fmt.Println("exit")
}

func wrongAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		go func(i int) {
			wg.Add(1)

			defer wg.Done()

			fmt.Println(i + 1)
		}(i)

	}

	wg.Wait()
	fmt.Println("exit")
}

func writeWithoutConcurrent() {
	start := time.Now()
	var counter int

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithoutMutex() {
	start := time.Now()
	var counter int
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++ // BUG: Несколько ГОРУТИН могут обратиться одновременно к ресурсу counter = counter + 1 // 555 + 1 = 556 // 555 + 1 = 556. Теряются операции
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func writeWithMutex() { // NOTE: Mutex исключает ситуацию, когда несколько операций ГОРУТИН могут обратиться одновременно к ресурсу. Не теряются операции.
	start := time.Now()
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			mu.Lock()   // NOTE: Заблокировали одновременный доступ ГОРУТИН
			counter++   // NOTE: Одна горутина получает доступ
			mu.Unlock() // NOTE: Разблокирован доступ для других горутин
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	start := time.Now()
	var (
		counter int
		wg      sync.WaitGroup
		mu      sync.RWMutex
	)

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.RLock() // NOTE: заблокировано на Чтение

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.RUnlock() // NOTE: Unlock fo READ
		}()

		go func() {
			defer wg.Done()

			mu.Lock() // NOTE: LOCK fo WRITE

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock() // NOTE: Unlock fo write
		}()
	}

	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
