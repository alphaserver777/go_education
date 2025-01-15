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
	// wrongAdd()

	// Mutex
	// writeWithoutConcurrent()
	// writeWithoutMutex()
	// writeWithMutex()
	// readWithMutex()
	readWithRWMutex()
}

func withoutWait() {
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	fmt.Println("exit")
}

func withWait() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			//defer wg.Done()
			fmt.Println(i + 1)
			wg.Done()
		}(i)

	}

	wg.Wait()
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
