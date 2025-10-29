// Программа для управления небольшим складом товаров
package main

import (
	"fmt"
	"time"
)

// Глобальные константы для категорий товаров
const (
	CategoryElectronics = "Электроника"
	CategoryFood        = "Продукты"
	CategoryClothes     = "Одежда"
	MaxItems            = 100 // Максимальное количество товаров на складе
)

var totalItems int
var nextItemID int64 = 12346

// Структура товаров
type Item struct {
	ID                       int64
	newItemID                int64
	Name                     string
	Category                 string
	Quantity                 int
	Price                    float64
	Discount                 float64
	IsAvailable              bool
	Color                    string
	Weight                   float32
	DateAdded                time.Time
	minQuantity, maxQuantity int
}

func main() {

	fmt.Printf("=== Информация о товаре ===\n")

	_, newItem := newItemAdd() // Получаем созданный новый товар
	displayItemInfo(&newItem)  // Передаём указатель на новый товар

	_, New1Item := New1ItemAdd()
	displayItemInfo(&New1Item)
}

// Функция для отображения информации о товаре
func displayItemInfo(item *Item) {
	fmt.Printf("ID: %d\n", item.ID)
	fmt.Printf("Название: %s\n", item.Name)
	fmt.Printf("Категория: %s\n", item.Category)
	fmt.Printf("Цвет: %s\n", item.Color)
	fmt.Printf("Вес товара: %.1f кг\n", item.Weight)
	fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", item.minQuantity, item.maxQuantity)

	// Расчет % максимального количества на складе
	percentInStock := float64(item.Quantity) / float64(MaxItems) * 100

	fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", percentInStock)
	fmt.Printf("Дата добавления: %s\n", item.DateAdded.Format("2006-01-02 15:04:05"))
	fmt.Printf("Количество на складе: %d\n", item.Quantity)
	fmt.Printf("Цена: %.2f руб.\n", item.Price)

	// Проверка наличия скидки
	if item.Discount < 0 || item.Discount > 100 {
		fmt.Println("Скидка не распростроняется или указана неверно")
	} else {
		// Расчет скидки
		discountedPrice := calculateDiscount(item.Price, item.Discount)
		fmt.Printf("Цена с учетом скидки: %.2f руб.\n", discountedPrice)
	}

	// Проверка наличия на складе
	if item.IsAvailable {
		fmt.Println("Статус: В наличии")
	} else {
		fmt.Println("Статус: Нет в наличии")
	}

	// Обновление общего количества товаров
	updateTotalItems(item.Quantity)
	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)
}

// Функция расчета скидки
func calculateDiscount(Price float64, Discount float64) float64 {
	return Price * ((100 - Discount) / 100)
}

// Функция для обновления общего количества товаров
func updateTotalItems(Quantity int) {
	totalItems += Quantity

	// Проверка на превышение максимального количества
	if totalItems > MaxItems {
		fmt.Printf("Предупреждение: Превышено максимальное количество товаров на складе на: %d\n", totalItems-MaxItems) // Вывод разницы
		totalItems = MaxItems
	}
}

// Функция добавления нового товара и возврата ID
func newItemAdd() (int64, Item) {
	newItem := Item{
		ID:          nextItemID,
		Name:        "Планшет",
		Category:    CategoryElectronics,
		Quantity:    44,
		Price:       39900.59,
		Discount:    101,
		IsAvailable: true,
		Color:       "Розовый",
		Weight:      0.6,
		DateAdded:   time.Now(),
		minQuantity: 2,
		maxQuantity: 10,
	}

	newItemID := nextItemID // Сохраняем текущий ID
	nextItemID++

	return newItemID, newItem
}

func New1ItemAdd() (int64, Item) {
	newItem := Item{
		ID:          nextItemID,
		Name:        "Молоко",
		Category:    CategoryFood,
		Quantity:    90,
		Price:       100.00,
		Discount:    10,
		IsAvailable: true,
		Color:       "Белый",
		Weight:      1,
		DateAdded:   time.Now(),
		minQuantity: 10,
		maxQuantity: 100,
	}
	newItemID := nextItemID
	nextItemID++

	return newItemID, newItem
}
