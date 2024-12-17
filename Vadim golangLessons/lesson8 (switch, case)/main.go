package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	min = 1
	max = 6 // Включает значение 5
)

func main() {
	// Инициализация счётчика для значений
	counts := make(map[int]int)

	// Количество итераций
	const iterations = 1000

	for i := 0; i < iterations; i++ {
		value := randomize()
		counts[value]++
	}

	// Вывод подсчёта значений
	fmt.Println("Распределение значений:")
	for i := min; i < max; i++ {
		fmt.Printf("%d: %d раз(а)\n", i, counts[i])
	}

	// Построение графика
	err := createBarChart(counts)
	if err != nil {
		fmt.Println("Ошибка при создании графика:", err)
	}
}

func randomize() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func createBarChart(counts map[int]int) error {
	// Преобразуем данные в формат, подходящий для построения графика
	values := make(plotter.Values, max-min)
	for i := min; i < max; i++ {
		values[i-min] = float64(counts[i])
	}

	// Создаём новый график
	p := plot.New()
	p.Title.Text = "Распределение случайных значений"
	p.Y.Label.Text = "Частота"
	p.X.Label.Text = "Значение"

	// Создаём столбцы
	bar, err := plotter.NewBarChart(values, vg.Points(20))
	if err != nil {
		return err
	}
	bar.LineStyle.Width = vg.Length(0) // Убираем линии вокруг столбцов
	bar.Color = plotter.DefaultColors[0]

	// Настраиваем оси
	p.Add(bar)
	p.NominalX("1", "2", "3", "4", "5")

	// Сохраняем график в файл
	return p.Save(8*vg.Inch, 4*vg.Inch, "distribution.png")
}
