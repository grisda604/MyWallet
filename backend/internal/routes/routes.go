package routes

import (
	"github.com/Mywallet/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

// ฟังก์ชันแยกสำหรับจัดการ Route ให้เป็นระเบียบ
func SetupRoutes(app *fiber.App) {
	// Group API ทั้งหมดไว้ใต้ /api
	api := app.Group("/api")

	// --- Wallet Routes ---
	api.Get("/wallets", handlers.GetWallets)
	api.Post("/wallets", handlers.CreateWallet)
	api.Put("/wallets/:id", handlers.UpdateWallet)
	api.Delete("/wallets/:id", handlers.DeleteWallet)

	// --- Transaction Routes ---
	api.Get("/transactions", handlers.GetTransactions)
	api.Post("/transactions", handlers.CreateTransaction)
	api.Put("/transactions/:id", handlers.UpdateTransaction)
	api.Delete("/transactions/:id", handlers.DeleteTransaction)

	// --- Category Routes ---
	api.Get("/categories", handlers.GetCategories)
	api.Post("/categories", handlers.CreateCategory)
	api.Put("/categories/:id", handlers.UpdateCategory)
	api.Delete("/categories/:id", handlers.DeleteCategory)

	// --- Savings Goal Routes ---
	api.Get("/savings-goals", handlers.GetSavingGoals)
	api.Get("/savings-goals/:id", handlers.GetSavingGoalByID)
	api.Post("/savings-goals", handlers.CreateSavingGoal)
	api.Put("/savings-goals/:id", handlers.UpdateSavingGoal)
	api.Delete("/savings-goals/:id", handlers.DeleteSavingGoal)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "message": "Server is running"})
	})
}
