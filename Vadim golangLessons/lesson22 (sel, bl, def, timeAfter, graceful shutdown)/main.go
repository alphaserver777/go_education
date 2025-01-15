package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//baseSelect()
	gracefulShutdown()
}

func baseSelect() {
	bufferedChan := make(chan string, 3)
	bufferedChan <- "first"

	select {
	case str := <-bufferedChan:
		fmt.Println("read", str)
	case bufferedChan <- "second":
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

	unbufChan := make(chan int)

	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()

	select {
	case bufferedChan <- "third": // NOTE: Приоритетность операций внутри select
		fmt.Println("unblocking writing")
	case val := <-unbufChan:
		fmt.Println("blocking reading", val)
	case <-time.After(time.Millisecond * 1500):
		fmt.Println("time`s up")
	default:
		fmt.Println("default case")
	}

	resultChan := make(chan int)
	timer := time.After(time.Second) // timer outside loop ВАЖНО!

	go func() { // Горутина, которая записывает в канал 1000 записей поочереди пока не закончится таймер.
		defer close(resultChan) // По завершению цила - закрываем канал

		for i := 1; i <= 1000; i++ {
			select { // Если запускаем select в цикле, то таймер выводим за цикл! Иначе каждую итерацию будем обновлять таймер и не достигнем конца таймера

			case <-timer: //NOTE: (select  выбирает между ожиданием получения результата операции (к примеру - выполнение цикла) и истечением таймера. Если таймер срабатывает раньше, чем завершается операция (работа цикла), то выводится сообщение о таймауте.)
				fmt.Println("time`s up")
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

// gracefulShutdown функция реализует механизм плавного завершения программы
func gracefulShutdown() {
	// 1. Создаем канал sigChan для приема сигналов операционной системы.
	//    - Канал имеет буфер размером 1, чтобы избежать блокировки при отправке сигнала.
	sigChan := make(chan os.Signal, 1)

	// 2. Указываем, какие сигналы операционной системы мы хотим перехватывать.
	//    - syscall.SIGINT - сигнал прерывания (Ctrl+C).
	//    - syscall.SIGTERM - сигнал завершения программы.
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 3. Создаем таймер, который отправит сигнал через 10 секунд.
	timer := time.After(10 * time.Second)

	// 4. Используем select для ожидания либо сигнала от таймера, либо сигнала от операционной системы.
	select {
	// 5. Если таймер сработал, выводим сообщение "time`s up" и завершаем работу функции.
	case <-timer:
		fmt.Println("time`s up")
		return
		// 6. Если получен сигнал от операционной системы, выводим сообщение о полученном сигнале и завершаем работу функции.
	case sig := <-sigChan:
		fmt.Println("Stopped by signal:", sig)
		return
	}
}
