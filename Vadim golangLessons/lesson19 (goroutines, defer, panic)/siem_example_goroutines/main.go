package main

import (
	"fmt"
	"time"
)

// Функция для имитации сбора логов с сервера
func collectLogs(server string, logs chan string) {
	for i := 1; i <= 3; i++ { // Имитируем сбор 3 логов
		log := fmt.Sprintf("Лог %d от %s", i, server)
		logs <- log                        // Отправляем лог в канал
		time.Sleep(500 * time.Millisecond) // Имитация задержки при сборе логов
	}
	close(logs) // Закрываем канал после завершения работы
}

func main() {
	servers := []string{"Server1", "Server2", "Server3"} // Список серверов
	var allLogs []string                                 // Срез для хранения всех логов

	for _, server := range servers {
		logs := make(chan string) // Создаем канал для каждого сервера

		// Запускаем горутину для сбора логов с сервера
		go collectLogs(server, logs)

		// Собираем логи из канала
		for log := range logs {
			allLogs = append(allLogs, log) // Добавляем лог в общий срез
		}
	}

	// Выводим все собранные логи
	fmt.Println("Все собранные логи:")
	for _, log := range allLogs {
		fmt.Println(log)
	}
}
