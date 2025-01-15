package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	baseKnowledge()
	//workerPool()
}

func baseKnowledge() {
	// 1. Создаем пустой контекст (context.Background).
	//    - Используется как базовый контекст для всей цепочки вызовов.
	//    - Не имеет значений, таймаута или сигнала отмены.
	ctx := context.Background()
	fmt.Println("context.Background:", ctx)

	// 2. Создаем "контекст-заглушку" (context.TODO).
	//    - Аналогичен context.Background(), но используется как "заглушка" (todo) на случай,
	//    - когда мы временно не знаем, какой контекст использовать.
	//    - Рекомендуется заменять его на более подходящий контекст по мере разработки.
	toDo := context.TODO()
	fmt.Println("context.TODO:", toDo)

	// 3. Создаем новый контекст на основе ctx и добавляем в него значение.
	//    -  Используем context.WithValue для передачи данных (например, идентификатора запроса) между функциями.
	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println("context.WithValue, value:", withValue.Value("name"))

	// 4. Создаем новый контекст с сигналом отмены.
	//    - context.WithCancel возвращает новый контекст и функцию cancel(), которую можно использовать
	//    - для отмены контекста и всех его дочерних контекстов.
	withCancel, cancel := context.WithCancel(ctx)
	//  Выводим ошибку контекста, на данном этапе она равна nil
	fmt.Println("context.WithCancel, error до вызова cancel():", withCancel.Err())
	// Вызываем функцию отмены, которая передает сигнал отмены
	cancel()
	// Выводим ошибку контекста, она должна быть отличной от nil, после отмены
	fmt.Println("context.WithCancel, error после вызова cancel():", withCancel.Err())

	// 5. Создаем новый контекст с крайним сроком (deadline).
	//    - context.WithDeadline создает контекст, который будет автоматически отменен через 3 секунды.
	//    - Функция cancel, возращаемая из WithDeadline, позволяет отменить контекст вручную до истечения времени
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	//  defer cancel() - гарантирует отмену контекста после выполнения кода
	defer cancel()
	// Выводим время, до которого контекст будет считаться не отменённым
	deadline, ok := withDeadline.Deadline()
	fmt.Println("context.WithDeadline, deadline:", deadline, "ok:", ok)
	//  Выводим ошибку контекста, на данном этапе она равна nil
	fmt.Println("context.WithDeadline, error до истечения deadline:", withDeadline.Err())

	//  Пытаемся получить данные из канала .Done. Программа блокируется до тех пор, пока не истечёт time.Second*3 или не будет вызвана cancel()
	fmt.Println("context.WithDeadline, получение данных из .Done() :", <-withDeadline.Done())
	//  Вызов канала .Done() после его истечения не заблокирует программу, а вернёт пустую структуру
	fmt.Println("context.WithDeadline, получение данных из .Done() после таймаута:", withDeadline.Err())

	// 6. Создаем новый контекст с таймаутом.
	//    - context.WithTimeout создает контекст, который будет автоматически отменен через 2 секунды.
	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	// гарантирует отмену контекста после выполнения кода
	defer cancel()
	// Выводим канал, из которого можно будет получить данные в случае отмены
	fmt.Println("context.WithTimeout, channel from .Done():", withTimeout.Done())

}
func workerPool() {

	ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(context.Background(), time.Millisecond*20)
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			/*if i == 500 {
				cancel()
			}*/
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int
	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)
}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}
