package actioninfo

import "fmt"

// DataParser содержит методы для парсинга данных и вывода информации об активности.
type DataParser interface {
	// Parse обрабатывает строку с данными и заполняет структуру, реализующую интерфейс DataParser.
	// Формат входных данных зависит от конкретной структуры (например, Training или DaySteps).
	Parse(datastring string) (err error)

	// ActionInfo формирует и возвращает строку с данными об активности.
	// Формат выходных данных зависит от конкретной структуры (например, Training или DaySteps).
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных о тренировках или прогулках.
//
// Параметры:
//
// dataset []string — слайс строк с данными об активности.
// dp DataParser — объект, реализующий интерфейс DataParser (например, Training или DaySteps).
func Info(dataset []string, dp DataParser) {
	if dp == nil {
		fmt.Println("Error: DataParser instance is nil")
		return
	}

	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			fmt.Println(err)
			continue
		}
		summary, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(summary)
	}
}
