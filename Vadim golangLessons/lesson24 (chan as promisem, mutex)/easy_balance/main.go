package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock() // Захват мьютекса перед изменением переменной balance
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock() // Освобождение мьютекса после изменения
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(100, &wg)
	go deposit(200, &wg)
	wg.Wait()
	fmt.Printf("New Balance %d\n", balance)
}
