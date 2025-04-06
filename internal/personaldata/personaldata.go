package personaldata

import "fmt"

// Personal содержит информацию о пользователе.
type Personal struct {
	Name   string  // имя пользователя.
	Weight float64 // вес пользователя в килограммах.
	Height float64 // рост пользователя в метрах.
}

// Print выводит данные пользователя: имя, вес и рост.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг\nРост: %.2f м\n", p.Name, p.Weight, p.Height)
}
