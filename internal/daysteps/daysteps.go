package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	if data == "" {
		log.Printf("Ошибка: пустая строка\n")
		return 0, 0, fmt.Errorf("пустая строка данных")
	}

	for _, ch := range data {
		if ch == ' ' {
			log.Printf("Ошибка: содержит пробел\n")
			return 0, 0, fmt.Errorf("Недопустимые пробелы в данных")
		}

	}

	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		log.Printf("Ошибка: некорректная длинна данных\n")
		return 0, 0, fmt.Errorf("Некорректная длинна данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Некорректное количество шагов")
	}

	if steps <= 0 {
		log.Printf("Ошибка: некорректное количество шагов\n")
		return 0, 0, fmt.Errorf("Количество шагов должно быть положительным")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Некорретный парсинг продолжительности прогулки")
	}

	if duration <= 0 {
		log.Printf("Ошибка: некорректный формат шагов\n")
		return 0, 0, fmt.Errorf("Продолжительность прогулки должна быть положительной")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка: при парсинге данных:", err)
		return ""
	}

	distanceMeters := float64(steps) * stepLength

	distanceKilometers := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка: парсинге данных:", err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		steps, distanceKilometers, calories)
}
