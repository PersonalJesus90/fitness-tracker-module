package daysteps

import (
	"errors"
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

	ErrInvalidFormat := errors.New("invalid string format")

	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		log.Printf("error: incorrect length of data\n")
		return 0, 0, ErrInvalidFormat
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, ErrInvalidFormat
	}

	if steps <= 0 {
		log.Printf("error: incorrect number of steps\n")
		return 0, 0, ErrInvalidFormat
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, ErrInvalidFormat
	}

	if duration <= 0 {
		log.Printf("error: incorrect step format\n")
		return 0, 0, ErrInvalidFormat
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("error parsing data:", err)
		return ""
	}

	distanceMeters := float64(steps) * stepLength

	distanceKilometers := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("error parsing data:", err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		steps, distanceKilometers, calories)
}
