package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func Distance(steps int) float64 {
	return (float64(steps) * lenStep) / mInKm
}

// МeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps) / duration.Hours()
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес(кг.) пользователя.
// height float64 — рост(м.) пользователя.
// duration time.Duration — длительность тренировки.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, fmt.Errorf("height and weight must be greater than zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than zero")
	}
	speed := MeanSpeed(steps, duration)
	calories := ((walkingCaloriesWeightMultiplier * weight) +
		(speed*speed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
	return calories, nil
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес(кг.) пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than zero")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be greater than zero")
	}
	speed := MeanSpeed(steps, duration)
	calories := ((runningCaloriesMeanSpeedMultiplier * speed) - runningCaloriesMeanSpeedShift) * weight
	return calories, nil
}
