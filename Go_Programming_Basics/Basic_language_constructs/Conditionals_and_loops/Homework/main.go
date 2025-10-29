package main

import (
  "fmt"
)

func main() {
  for i := 0; i != 5; i++ {
    var age int

    // Запросить озраст пользователя
    fmt.Printf("Пожалуйста, введите ваш возраст: ")
    fmt.Scan(&age)

    // Определяем статус в зависимости от возраста

    if age < 0 {
      fmt.Println("Возраст не может быть отрицательным")
    } else if age < 18 {
      fmt.Println("Ребенок")
    } else if age >= 18 && age <= 65 {
      fmt.Println("Взрослый")
    } else {
      fmt.Println("Пенсионер")
    }
  }
}
