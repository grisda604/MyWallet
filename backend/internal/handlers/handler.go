package handlers

import (
	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/models"
	"github.com/Mywallet/internal/services"
	"github.com/gofiber/fiber/v2"
)

func GetTransactions(c *fiber.Ctx) error {
	var transactions []models.Transaction
	database.DB.Preload("Wallet").Preload("Category").Preload("TargetWallet").Order("date desc").Find(&transactions)
	return c.JSON(fiber.Map{"transactions": transactions})
}

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.CreateTransaction(transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Transaction created successfully", "data": transaction})
}

func UpdateTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	oldTx := new(models.Transaction)
	if err := database.DB.First(&oldTx, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaction not found"})
	}

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.UpdateTransaction(oldTx, transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Transaction updated successfully", "data": transaction})
}

func DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	transaction := new(models.Transaction)
	if err := database.DB.First(&transaction, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaction not found"})
	}

	err := services.DeleteTransaction(transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Transaction deleted successfully"})
}

// ==================== Wallet Handlers ====================

func GetWallets(c *fiber.Ctx) error {
	wallets, err := services.GetWallers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(wallets)
}

func CreateWallet(c *fiber.Ctx) error {
	wallet := new(models.Wallet)
	if err := c.BodyParser(wallet); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	// เรียก service layer ที่มี validation
	err := services.CreateWallet(wallet)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Wallet created successfully", "data": wallet})
}

func UpdateWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	updateData := new(models.Wallet)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	// เรียก service layer ที่มี validation
	err := services.ModifyWallet(id, updateData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// ดึงข้อมูล wallet ที่อัพเดทแล้วกลับมา
	wallet, err := services.GetWalletByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Wallet updated successfully", "data": wallet})
}

func DeleteWallet(c *fiber.Ctx) error {
	id := c.Params("id")

	// ตรวจสอบว่า wallet มีอยู่จริง
	_, err := services.GetWalletByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Wallet not found"})
	}

	// เรียก service layer เพื่อลบ
	err = services.DeleteWallet(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Wallet deleted successfully"})
}

// ==================== Category Handlers ====================

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.CreateCategory(category)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Category created successfully", "data": category})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	updateData := new(models.Category)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	// ตรวจสอบว่า category มีอยู่จริง
	category := new(models.Category)
	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}

	// อัพเดทข้อมูล
	if updateData.Name != "" {
		category.Name = updateData.Name
	}
	if updateData.Type != "" {
		category.Type = updateData.Type
	}
	if updateData.Color != "" {
		category.Color = updateData.Color
	}
	if updateData.Icon != "" {
		category.Icon = updateData.Icon
	}

	if err := database.DB.Save(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category updated successfully", "data": category})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	// ตรวจสอบว่า category มีอยู่จริง
	category := new(models.Category)
	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}

	err := services.DeleteCategory(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}

func GetSavingGoals(c *fiber.Ctx) error {
	savings, err := services.GetSavingsGoals()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(savings)
}

func GetSavingGoalByID(c *fiber.Ctx) error {
	id := c.Params("id")

	savings, err := services.GetSavingsGoalByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(savings)
}

func CreateSavingGoal(c *fiber.Ctx) error {
	Savings := new(models.SavingsGoal)
	if err := c.BodyParser(Savings); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.CreateSavingsGoal(Savings)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Saving goal created successfully", "data": Savings})
}

func UpdateSavingGoal(c *fiber.Ctx) error {
	id := c.Params("id")
	updateData := new(models.SavingsGoal)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.UpdateSavingsGoal(id, *updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	savings, err := services.GetSavingsGoalByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Saving goal updated successfully", "data": savings})
}

func DeleteSavingGoal(c *fiber.Ctx) error {
	id := c.Params("id")

	// ตรวจสอบว่า savings goal มีอยู่จริง
	savings := new(models.SavingsGoal)
	if err := database.DB.First(&savings, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Savings goal not found"})
	}

	err := services.DeleteSavingGoal(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Savings goal deleted successfully"})
}
