package service

import (
	"prtask1/internal/repo"
	"prtask1/pkg/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Слой бизнес-логики. Тут должна быть основная логика сервиса

// Service - интерфейс для бизнес-логики
type Service interface {
	CreateTask(c *fiber.Ctx) error
	GetAllTasks(c *fiber.Ctx) error
	GetTaskByID(c *fiber.Ctx) error
	UpdateTask(c *fiber.Ctx) error
	DeleteTask(c *fiber.Ctx) error
}

type service struct {
	storage *repo.MemoryStorage
}

// NewService - конструктор сервиса
func NewService(storage *repo.MemoryStorage) *service {
	return &service{
		storage: storage,
	}
}

// Создать задачу
func (s *service) CreateTask(c *fiber.Ctx) error {
	var req repo.Task

	// Парсинг JSON
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request body",
		})
	}

	// Валидация
	if err := validator.Validate(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Создание задачи через storage
	task := &repo.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	createdTask := s.storage.Create(task)

	// Формирование ответа
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   createdTask,
	})
}

// Получить все задачи
func (s *service) GetAllTasks(c *fiber.Ctx) error {
	tasks := s.storage.GetAll()

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   tasks,
	})
}

// Получить задачу по ID
func (s *service) GetTaskByID(c *fiber.Ctx) error {
	// Извлечение ID из URL
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid task ID",
		})
	}

	// Получение задачи из storage
	task, err := s.storage.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Task not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   task,
	})
}

// Обновить задачу
func (s *service) UpdateTask(c *fiber.Ctx) error {
	// Извлечение ID из URL
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid task ID",
		})
	}

	var req repo.Task

	// Парсинг JSON
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request body",
		})
	}

	// Валидация
	if err := validator.Validate(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// Обновление через storage
	updatedTask := &repo.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	task, err := s.storage.Update(id, updatedTask)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Task not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   task,
	})
}

// Удалить задачу
func (s *service) DeleteTask(c *fiber.Ctx) error {
	// Извлечение ID из URL
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid task ID",
		})
	}

	// Удаление через storage
	err = s.storage.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Task not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Task deleted successfully",
	})
}
