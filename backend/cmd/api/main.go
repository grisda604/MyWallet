package main

import (
	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. เชื่อมต่อ Database (และ Auto Migrate ตามที่เราเขียนไว้)
	database.ConnectDB()

	// 2. สร้าง Fiber App
	app := fiber.New(fiber.Config{
		AppName: "Expense Tracker API",
	})

	// 3. Middlewares
	// Logger: ช่วยแสดง Log เวลายิง API (เช่น GET /api/wallets 200 OK)
	app.Use(logger.New())

	// CORS: สำคัญมาก! อนุญาตให้ Nuxt (Port 3000) ยิงเข้ามาได้
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:8080, http://localhost:8081",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// 4. Setup Routes
	routes.SetupRoutes(app)

	// 5. Start Server ที่ Port 8081
	log.Info("Server is running on http://localhost:8081")
	log.Fatal(app.Listen(":8081"))
}
