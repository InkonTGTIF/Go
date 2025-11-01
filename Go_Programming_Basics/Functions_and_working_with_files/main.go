package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Читает данные из файла "input.txt", обрабатывает их
func ReadProcessWrite(inputPath, outputPath string, processFunc func(string) (string, error)) error {
	// Чтение файла
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла %s: %w", inputPath, err)
	}

	// Обработка данных
	processedData, err := processFunc(string(inputData))
	if err != nil {
		return fmt.Errorf("Ошибка обработки данных: %w", err)
	}

	// Запись данных в файл
	err = os.WriteFile(outputPath, []byte(processedData), 0644)
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл %s: %w", outputPath, err)
	}

	return nil
}

func main() {
	// Преобразование в верхний регистр
	err := ReadProcessWrite("input.txt", "output.txt", func(data string) (string, error) {
		return strings.ToUpper(data), nil
	})

	// Обработка ошибок
	if err != nil {
		// Логируем и завершение с ошибкой
		log.Fatalf("Ошибка: %w", err)
	} else {
		fmt.Println("Файл обработан успешно.")
	}
}
