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
	unpacked := strings.Split(data, ",")
	if len(unpacked) != 2 {
		return 0, 0, fmt.Errorf("unproper string")
	}
	steps, err := strconv.Atoi(unpacked[0])
	if err != nil {
		return 0, 0, err
	}
	if steps < 1 {
		return 0, 0, fmt.Errorf("steps is not natural number")
	}
	duration, err := time.ParseDuration(unpacked[1])
	if err != nil {
		return 0, 0, fmt.Errorf("duration is not available to parsing")
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("duration in not positive")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, time, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	distance := stepLength * float64(steps)
	dist_km := distance / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, time)
	if err != nil {
		log.Println(err)
		return ""
	}
	actionInfo := fmt.Sprintf("Количество шагов: %d.\n", steps)
	actionInfo += fmt.Sprintf("Дистанция составила %.2f км.\n", dist_km)
	actionInfo += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)
	return actionInfo
}
