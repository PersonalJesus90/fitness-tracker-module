package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {

	if data == "" {
		return 0, "", 0, fmt.Errorf("empty row")
	}

	for _, ch := range data {
		if ch == ' ' {
			return 0, "", 0, fmt.Errorf("invalid spaces")
		}

	}

	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("incorrect time format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("step parsing error")
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("the number of steps must be positive")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("duration parsing error")
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("the duration must be positive")
	}

	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceInMeters := float64(steps) * stepLength
	return distanceInMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	var calories float64
	var errCalories error

	switch activityType {
	case "Walk":
		calories, errCalories = WalkingSpentCalories(steps, weight, height, duration)
	case "Run":
		calories, errCalories = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("unknown type of training")
	}

	if errCalories != nil {
		return "", errCalories
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	durationHours := duration.Hours()

	result := fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		activityType, durationHours, dist, speed, calories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("the values: steps, weight, height, duration, must be positive")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * avgSpeed * durationInMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("the values: steps, weight, height, duration, must be positive")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * avgSpeed * durationInMinutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
