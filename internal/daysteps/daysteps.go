package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// DaySteps содержит информацию о дневных прогулках пользователя.
type DaySteps struct {
	personaldata.Personal               // данные пользователя (имя, вес, рост).
	Steps                 int           // количество шагов.
	Duration              time.Duration // длительность прогулки.
}

// Parse парсит строку с данными формата "678,0h50m"
// и записывает данные в соответствующие поля структуры DaySteps.
//
// Параметры:
//
// datastring string — строка, содержащая количество шагов и продолжительность, разделённые запятой.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format: expected 2 values, got %d, "+
			"data: %q", len(parts), datastring)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	ds.Duration = duration
	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке:
// количество шагов, дистанцию (км.) и количество сожжённых калорий.
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration must be greater than zero")
	}
	distance := spentenergy.Distance(ds.Steps)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	summary := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.", ds.Steps, distance, calories)
	return summary, nil
}
