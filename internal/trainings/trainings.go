package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// Training содержит информацию о тренировке пользователя.
type Training struct {
	personaldata.Personal               // данные пользователя (имя, вес, рост).
	Steps                 int           // количество шагов за тренировку.
	TrainingType          string        // тип тренировки (бег или ходьба).
	Duration              time.Duration // длительность тренировки.
}

// Parse парсит строку с данными формата "3456,Ходьба,3h00m"
// и записывает данные в соответствующие поля структуры Training.
//
// Параметры:
//
// datastring string — строка, содержащая количество шагов, тип тренировки и продолжительность, разделённые запятой.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format: expected 3 values, got %d, "+
			"data: %q", len(parts), datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	if parts[1] != "Ходьба" && parts[1] != "Бег" {
		return fmt.Errorf("unknown activity type: expected 'Ходьба' or 'Бег', got %q", parts[1])
	}
	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	t.Duration = duration
	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке:
// тип тренировки, длительность (ч.), дистанцию (км.), среднюю скорость (км/ч) и количество сожжённых калорий.
func (t Training) ActionInfo() (string, error) {
	if t.Duration <= 0 {
		return "", fmt.Errorf("duration must be greater than zero")
	}
	distance := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	default:
		return "неизвестный тип тренировки", fmt.Errorf("unknown training type")
	}
	if err != nil {
		return "", err
	}

	summary := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), distance, speed, calories)
	return summary, nil
}
