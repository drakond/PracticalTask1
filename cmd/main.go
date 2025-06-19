package main

import (
	"log"
	"prtask1/internal/api"
	"prtask1/internal/config"
	"prtask1/internal/repo"
	"prtask1/internal/service"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Инициализация зависимостей
	memoryStorage := repo.NewMemoryStorage()
	taskService := service.NewService(memoryStorage)

	// Настройка роутов
	app := api.SetupRoutes(taskService)

	// Запуск сервера
	log.Printf("🚀 %s starting on port %s", cfg.AppName, cfg.Port)
	app.Listen(cfg.Port)
}
