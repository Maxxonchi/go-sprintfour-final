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
	unpacked := strings.Split(data, ",")
	if len(unpacked) != 3 {
		return 0, "", 0, fmt.Errorf("incorrect data")
	}
	steps, err := strconv.Atoi(unpacked[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(unpacked[2])
	if err != nil {
		return 0, "", 0, err
	}
	return steps, unpacked[1], duration, nil
}

func distance(steps int, height float64) float64 {
	step_length := stepLengthCoefficient * height
	dist := float64(steps) * step_length
	dist_km := dist / mInKm
	return dist_km
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration == 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, train, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	var calories float64
	switch train {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	trainingInfo := fmt.Sprintf("Тип тренировки: %s\n", train)
	trainingInfo += fmt.Sprintf("Длительность: %.2f ч.\n", duration.Minutes()/minInH)
	trainingInfo += fmt.Sprintf("Дистанция: %.2f км.\n", dist)
	trainingInfo += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	trainingInfo += fmt.Sprintf("Сожгли калорий: %.2f", calories)
	return trainingInfo, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 1 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect values")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 1 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("incorrect values")
	}
	speed := meanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}
