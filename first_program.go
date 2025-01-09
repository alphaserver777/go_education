package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	workMinutes  = 25
	breakMinutes = 5
	currentMode  = ""
	timeLeft     = 0
	mutex        sync.Mutex
)

func main() {
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/progress", progressHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Таймер запущен!")
	go func() {
		startTimer(workMinutes, "Работа")
		startTimer(breakMinutes, "Отдых")
	}()
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	work := query.Get("work")
	breakTime := query.Get("break")

	if work != "" {
		if value, err := strconv.Atoi(work); err == nil {
			workMinutes = value
		}
	}

	if breakTime != "" {
		if value, err := strconv.Atoi(breakTime); err == nil {
			breakMinutes = value
		}
	}

	fmt.Fprintf(w, "Настройки обновлены: Работа = %d минут, Отдых = %d минут\n", workMinutes, breakMinutes)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Текущие настройки: Работа = %d минут, Отдых = %d минут\n", workMinutes, breakMinutes)
}

func progressHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	minutes := timeLeft / 60
	seconds := timeLeft % 60
	fmt.Fprintf(w, "Режим: %s\nОсталось времени: %02d:%02d\n", currentMode, minutes, seconds)
}

func startTimer(minutes int, mode string) {
	mutex.Lock()
	currentMode = mode
	timeLeft = minutes * 60
	mutex.Unlock()

	for timeLeft > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		timeLeft--
		mutex.Unlock()
	}

	mutex.Lock()
	currentMode = ""
	mutex.Unlock()

	fmt.Printf("%s завершен!%s\n", mode, spacesToClearLine())
}

func spacesToClearLine() string {
	return "                        " // Для очистки строки после \r
}
