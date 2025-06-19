package main

import (
	"prtask1/internal/api"
	"prtask1/internal/repo"
	"prtask1/internal/service"
)

func main() {
	// Инициализация зависимостей
	memoryStorage := repo.NewMemoryStorage()
	taskService := service.NewService(memoryStorage)

	// Настройка роутов
	app := api.SetupRoutes(taskService)

	// Запуск сервера
	app.Listen(":3000")
}
